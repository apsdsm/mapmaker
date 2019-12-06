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

type DungeonImporter struct {
	Dungeon         *placeholder.Dungeon
	Errors          *ErrorList
	annotationTypes map[string]string
}

type DungeonImporterConfig struct {
	Errors *ErrorList
}

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

func (i *DungeonImporter) Read(in string) error {
	var err error

	// get the meta and dungeon for this file as buffers
	metaBuffer, dungeonBuffer := getMetaAndDungeonBuffers(in)

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

// getMtaAndMapBuffers takes the path to a .map file and returns two buffers: one for the meta and one for the dungeon
func (i *DungeonImporter) getMetaAndDungeonBuffers(path string) (metaBuffer bytes.Buffer, dungeonBuffer bytes.Buffer) {

	// load map file
	path, err := filepath.Abs(path)
	data, err := os.Open(path)
	defer data.Close()

	if err != nil {
		panic("error wile opening file")
	}

	// load map file into scanner
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	// loop variables
	section := beforeMeta
	seperator := "---"
	newline := "\n"

	// .map files have the following layout
	//     ---                    <- beforeMeta
	//     yaml structure meta    <- insideMeta
	//     ---
	//     map structure          <- mapContent
	//
	// this loop will write lines to the correct buffer based on where the cursor is in the file.
	for scanner.Scan() {
		text := scanner.Text()

		if text == seperator && section == beforeMeta {
			section = insideMeta
		} else if text == seperator && section == insideMeta {
			section = mapContent
		} else if section == insideMeta {
			metaBuffer.WriteString(text + newline)
		} else if section == mapContent {
			dungeonBuffer.WriteString(text + newline)
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

	// copy cells from lines to grid
	// - converts the alignment of rows and cols (lines are ROWxCOL, but a grid is COLxROW)
	//   ```
	//   〈 a, b, c 〉
	//   〈 d, e, f 〉
	//   〈 g, h, r 〉
	//   ```
	//   will be transformed to:
	//   ```
	//     ︿  ︿  ︿
	//     a   b   c
	//     d   e   f
	//     g   h   r
	//     ﹀  ﹀  ﹀
	//   ```
	//
	// - pads empty cells to the right if line is shorter than map width
	// - checks for start cell and sets dungeon data appropriately
	// - sets a ' ' as the rune value for start cell and entity cells
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			// pad with empty space if this line is shorter than longest line
			if c > len(lines[r].Cells) {
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

// Each line of the dungeon can have an optional annotation (denoted by //) which contains information about
// the entities on that line. Annotations are attached left to right.
//
// ```
// #   m m   # // m:goblin m:orc
// ```
func (i *DungeonImporter) getPartsFromRawLine(text string) (mapLine placeholder.Line, mapAnnotations []placeholder.Annotation) {

	parts := strings.Split(text, "//")

	// if map data present
	if len(parts) > 0 {
		rawMapLine := strings.TrimRight(parts[0], " ")
		mapLine := placeholder.NewLine(len(rawMapLine))

		// convert the runes in the line into cells
		for i, r := range rawMapLine {
			mapLine.Cells[i] = placeholder.NewCellFromRune(r)
		}
	}

	// if annotations present
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
