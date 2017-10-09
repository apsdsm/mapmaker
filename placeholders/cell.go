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

package placeholders

// Cell is a placeholder for a single cell in a map
type Cell struct {
	// content of cell as read from file
	Rune rune

	// type of cell
	Type string

	// link from annotation
	Link string

	// true if this cell was annotated
	Annotated bool

	// location coordinates for tile
	X, Y int
}

// EmptyCell returns a blank empty cell
func EmptyCell() *Cell {
	c := Cell{}
	return &c
}

// NewCellFromRune makes a new cell from the given rune
func NewCellFromRune(r rune) *Cell {
	c := Cell{
		Rune: r,
	}

	return &c
}

// IsEntity returns true if the cell contains an entity rune
func (c *Cell) IsEntity() bool {
	for _, t := range []rune{'m', 'd', 'w', 's'} {
		if c.Rune == t {
			return true
		}
	}
	return false
}

// IsStart returns true if the cell contains a start rune
func (c *Cell) IsStart() bool {
	return c.Rune == 's'
}
