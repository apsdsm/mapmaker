package file

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/apsdsm/mapmaker/formats/placeholder"
)

const (
	beforeMeta = iota
	insideMeta
	insideDungeon
)

// lineParts are the two parts of a line of the dungeon after it is split into map data and annotation data
type lineParts []string

// A DungeonImporter creates a placeholder.Dungeon from a .map file.
type DungeonImporter struct {
	Dungeon         *placeholder.Dungeon
	Errors          *ErrorList
	annotationTypes map[string]string
}

// A DungeonImporterConfig contains the config data for a DungeonImporter.
type DungeonImporterConfig struct {
	Errors *ErrorList
}

// NewDungeonImporter will create and initialize a new DungeonImporter.
func NewDungeonImporter(c DungeonImporterConfig) *DungeonImporter {
	i := DungeonImporter{
		Dungeon: &placeholder.Dungeon{},
		Errors:  c.Errors,
		annotationTypes: map[string]string{
			"m": "mob",
			"d": "door",
			"w": "waypoint",
			"s": "start",
			"i": "item",
		},
	}

	return &i
}

// Read will load a raw dungeon file, and convert it into a placeholder.Dungeon object. It will record any errors
// in the DungeonImporter's Error list.
func (i *DungeonImporter) Read(in string) error {
	var err error

	// get the meta and dungeon for this file as buffers
	metaBuffer, dungeonBuffer := i.getMetaAndDungeonBuffers(in)

	// parse meta
	if err := i.parseMetaBuffer(metaBuffer); err != nil {
		return err
	}

	// parse dungeon
	if err = i.parseDungeonBuffer(dungeonBuffer); err != nil {
		return err
	}
	return nil
}

// takes the location of a .map file, and splits it into two buffers - one for th meta data and one for the dungeon.
func (i *DungeonImporter) getMetaAndDungeonBuffers(path string) (metaBuffer bytes.Buffer, dungeonBuffer bytes.Buffer) {
	// open file and load into scanner
	path, err := filepath.Abs(path)
	data, err := os.Open(path)

	if err != nil {
		panic("error wile opening file")
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	// scan through the text file, writing lines to the appropriate buffer:
	//
	// .map files have the following layout
	//                            <- `beforeMeta`
	//     ---                    <- separator indicates change of section
	//     yaml structure meta    <- `insideMeta` (this content is written to the metaBuffer)
	//     ---                    <- separator indicates change of section
	//     map structure          <- `insideDungeon` (this content is written to the dungeonBuffer)
	//     ...                    <- continues to end of document

	section := beforeMeta

	for scanner.Scan() {
		text := scanner.Text()

		if text == "---" && section == beforeMeta {
			section = insideMeta
		} else if text == "---" && section == insideMeta {
			section = insideDungeon
		} else if section == insideMeta {
			metaBuffer.WriteString(text + "\n")
		} else if section == insideDungeon {
			dungeonBuffer.WriteString(text + "\n")
		}
	}

	return metaBuffer, dungeonBuffer
}

// parseMetaBuffer will marshal a yaml meta buffer into a placeholder object.
func (i *DungeonImporter) parseMetaBuffer(metaBuffer bytes.Buffer) error {
	return yaml.Unmarshal(metaBuffer.Bytes(), i.Dungeon)
}

// parseDungeonBuffer will parse a dungeon buffer into an array of placeholder lines. It also returns rows and cols in dungeon
func (i *DungeonImporter) parseDungeonBuffer(dungeonBuffer bytes.Buffer) error {

	// put map data into a scanner
	reader := bytes.NewReader(dungeonBuffer.Bytes())
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	// set up vars for the loop
	lines := make([]*placeholder.Line, 0, 64)
	rows := 0
	cols := 0

	// scan through the map, adding lines as we go
	for scanner.Scan() {
		rawLine := scanner.Text()

		// do not read any empty lines immediately after the meta block - this allows the map edtiro to insert some
		// padding into the map file if they want.
		if rows == 0 && rawLine == "" {
			continue
		}

		line, annotations := i.getPartsFromRawLine(rawLine)

		if line.Len() > cols {
			cols = line.Len()
		}

		rows++

		// Try find a match for each annotation. If there is no match for a specific annotation, add an error.
		// Annotations are attached left to right, and to the first matching, unannotated cell.
		for _, a := range annotations {
			for _, c := range line.Cells {

				// don't bother if already annotated
				if c.Annotated {
					continue
				}

				// if the rune for a cell matches the a target, and we have not yet annotated that cell
				if string(c.Rune) == a.Target {
					var ok bool

					if c.Type, ok = i.annotationTypes[a.Target]; !ok {
						panic("unknown annotation type")
					}

					c.Annotated = true
					c.Link = a.Link
					a.Assigned = true

					break
				}
			}
		}

		lines = append(lines, &line)
	}

	// allocate enough space for dungeon
	i.Dungeon.AllocateTiles(cols, rows)

	// copy cells from lines to grid:
	// - convert the alignment of rows and cols (lines are ROWxCOL, but a grid is COLxROW)
	//
	//   ```
	//   〈 a  b  c 〉  <- from this
	//   〈 d  e  f 〉
	//   〈 g  h  r 〉
	//
	//     ︿  ︿  ︿
	//     a   b   c   <- to this
	//     d   e   f
	//     g   h   r
	//     ﹀  ﹀  ﹀
	//   ```
	//
	// - pad empty cells to the right if line is shorter than map width
	// - check for start cell and sets dungeon data appropriately
	// - set ' ' as the rune value for start cell and entity cells
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			// pad with empty space if this line is shorter than longest line
			if c >= len(lines[r].Cells) {
				i.Dungeon.Grid[c][r] = placeholder.EmptyCell()
				continue
			}

			cell := lines[r].Cells[c]

			// if this is a start tile
			if cell.IsStart() {
				i.Dungeon.StartPosition = &placeholder.Position{X: c, Y: r}
				cell.Rune = ' '
			}

			// if this is an entity
			if cell.IsEntity() {
				cell.Rune = ' '
			}

			i.Dungeon.Grid[c][r] = cell
		}
	}

	return nil
}

// split a line of a dungeon into its layout and annotation parts. Each line of the dungeon can have an optional
// annotation (denoted by //) which contains information about the entities on that line. Annotations are attached
// left to right.
//
// ```
//     +- this is the monkey
//     | +-- this is the lion
//     | |
// #   m m  # // m:monkey m:lion
// ^^^^^^^^^^    ^^^^^^^^^^^^^^^
//   layout        annotations
// ```
func (i *DungeonImporter) getPartsFromRawLine(text string) (placeholder.Line, []placeholder.Annotation) {
	var mapAnnotations []placeholder.Annotation
	var mapLine placeholder.Line

	parts := strings.Split(text, "//")

	// if map data present, convert the runes on the line to cells
	if len(parts) > 0 {
		rawMapLine := strings.TrimRight(parts[0], " ")
		mapLine = placeholder.NewLine(len(rawMapLine))

		for i, r := range rawMapLine {
			mapLine.Cells[i] = placeholder.NewCellFromRune(r)
		}
	}

	// if annotations present, assign them to to the cells from the previous step
	if len(parts) > 1 {
		rawAnnotations := strings.Split(parts[1], " ")

		for _, r := range rawAnnotations {
			aParts := strings.Split(r, ":")

			if len(aParts) != 2 {
				continue
			}

			a := placeholder.Annotation{
				Target: aParts[0],
				Link:   aParts[1],
			}

			mapAnnotations = append(mapAnnotations, a)
		}
	}

	return mapLine, mapAnnotations
}
