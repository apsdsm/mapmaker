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
	"github.com/apsdsm/mapmaker/formats/placeholder"
)

// ValidateDungeon will run all map data through a set of validations to make sure nothing is broken
func ValidateDungeon(mapMeta *placeholder.Meta, levelData *placeholder.Map, entities *placeholder.EntityCollection) (errors []Error, warnings []Error) {
	errors = make([]Error, 0, 10)
	warnings = make([]Error, 0, 10)

	// check for name
	if mapMeta.Name == "" {
		warn := Error{
			Message:    "map has no name.",
			LineNumber: -1,
			IsWarning:  true,
		}
		errors = append(warnings, warn)
	}

	// check for start point
	if levelData.StartPosition == nil {
		err := Error{
			Message:    "map has no start position.",
			LineNumber: -1,
		}
		errors = append(errors, err)
	}

	// validate each cell
	for x := 0; x < levelData.Width; x++ {
		for y := 0; y < levelData.Height; y++ {
			err := validateCell(levelData.Grid[x][y], y+1, entities)

			if err != nil {
				if err.IsWarning {
					warnings = append(warnings, *err)
				} else {
					errors = append(errors, *err)
				}
			}
		}
	}

	return errors, warnings
}

// validate a single cell
func validateCell(cell *placeholder.Cell, line int, entities *placeholder.EntityCollection) *Error {
	if !cell.Annotated {
		return nil
	}

	if cell.Type == "mob" {
		if !entities.HasMob(cell.Link) {
			return &Error{
				Message:    cell.Link + " is not defined by any entity.",
				LineNumber: line,
			}
		}
	}

	if cell.Type == "door" {
		if !entities.HasDoor(cell.Link) {
			return &Error{
				Message:    cell.Link + " is not defined by any entity.",
				LineNumber: line,
			}
		}
	}

	return nil
}
