// Copyright 2017 Nick del Pozo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/apsdsm/mapmaker/formats/placeholder"
	yaml "gopkg.in/yaml.v2"
)

const (
	beforeMeta = iota
	insideMeta
	mapContent
)

// lineParts are the two parts of a line of the dungeon after it is split into map data and annotation data
type lineParts []string

// ImportDungeon will import the data from a .map file and store it in a placeholder
func ImportDungeon(path string) (*placeholder.Meta, *placeholder.Map) {

	// get the meta and dungeon for this file as buffers
	metaBuffer, dungeonBuffer := getMetaAndDungeonBuffers(path)

	// parse buffers
	meta := parseMetaBuffer(metaBuffer)
	dungeon := parseDungeonBuffer(dungeonBuffer)

	// return parsed placeholders
	return meta, dungeon
}

// getMtaAndMapBuffers takes the path to a .map file and returns two buffers: one for the meta and one for the dungeon
func getMetaAndDungeonBuffers(path string) (metaBuffer bytes.Buffer, dungeonBuffer bytes.Buffer) {

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
func parseMetaBuffer(metaBuffer bytes.Buffer) *placeholder.Meta {
	var meta placeholder.Meta

	err := yaml.Unmarshal(metaBuffer.Bytes(), &meta)

	if err != nil {
		panic("trouble unmarshaling map meta")
	}

	return &meta
}

// parseDungeonBuffer will parse a dungeon buffer into an array of placeholder lines. It also returns rows and cols in dungeon
func parseDungeonBuffer(dungeonBuffer bytes.Buffer) *placeholder.Map {

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

		// do not read empty lines immediately after the meta block
		if rows == 0 && rawLine == "" {
			continue
		}

		// only advance row count if there is content on row
		rows++

		// get the map data and annotation data from the raw line
		mapData, annotations := getPartsFromRawLine(rawLine)

		// if this line (annotations excluded) is the longest line we've encountered so far, make it the map width
		if len(mapData) > cols {
			cols = len(mapData)
		}

		// prepare a new line
		line := placeholder.NewLine(len(mapData))

		// for each cell in the map line, add them as runes to the line
		for i, r := range mapData {
			line.Cells[i] = placeholder.NewCellFromRune(r)
		}

		// for each annotation, parse them and add them to the line
		for _, annotation := range annotations {
			for _, c := range line.Cells {
				// if the rune for a cell matches the annotation target, and we have not yet annotated that cell
				if string(c.Rune) == annotation.Target && !c.Annotated {
					c.Annotated = true
					c.Type = getEntityType(annotation.Target)
					c.Link = annotation.Link
					break
				}
			}
		}

		lines = append(lines, &line)
	}

	// make placeholder dungeon
	dungeon := placeholder.NewMap(rows, cols)

	// copy cells from lines to grid
	// - converts the alignment of rows and cols (lines are ROWxCOL, but a grid is COLxROW)
	// - populates with empty cells where there is no cell data
	// - checks for start cell and sets dungeon data appropriately
	// - sets a ' ' as the rune value for start cell and entity cells
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j < len(lines[i].Cells) {
				thisCell := lines[i].Cells[j]

				// if this is a start tile
				if thisCell.IsStart() {
					dungeon.StartPosition = &placeholder.Position{X: j, Y: i}
					thisCell.Rune = ' '
				}

				// if this is an entity
				if thisCell.IsEntity() {
					thisCell.Rune = ' '
				}

				dungeon.Grid[j][i] = thisCell
			} else {
				// pad with empty space if this line is shorter than longest line
				dungeon.Grid[j][i] = placeholder.EmptyCell()
			}
		}
	}

	return dungeon
}

// Each line of the dungeon is split into map data, and an optional array of annotation for that line (split
// at a `//`). This function will return the
// ```
// #   m   # // m:goblin
// ```
//
func getPartsFromRawLine(text string) (mapData string, annotations []placeholder.Annotation) {
	parts := strings.Split(text, "//")

	if len(parts) > 0 {
		mapData = strings.TrimRight(parts[0], " ")
	}

	if len(parts) > 1 {
		annotations = makeAnnotationArray(parts[1])
	}

	return
}

// getAnnotations will parse and return annotation placeholders for each valid raw annotation in the array.
func makeAnnotationArray(annotationData string) []placeholder.Annotation {
	annotationData = strings.Trim(annotationData, " ")
	rawAnnotations := strings.Split(annotationData, " ")

	annotations := make([]placeholder.Annotation, 0, len(rawAnnotations))

	for _, rawAnnotation := range rawAnnotations {
		aParts := strings.Split(rawAnnotation, ":")

		if len(aParts) != 2 {
			continue
		}

		a := placeholder.Annotation{
			Target: aParts[0],
			Link:   aParts[1],
		}

		annotations = append(annotations, a)
	}

	return annotations
}

// getEntityType returns the appropriate entity type for the given key.
func getEntityType(key string) string {
	switch key {
	case "m":
		return "mob"
	case "d":
		return "door"
	case "w":
		return "waypoint"
	case "s":
		return "start"
	case "i":
		return "item"
	}

	panic(key + " is an unknown annotation target type")
}
