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

// A Dungeon contains map and entity data for a single dungeon
type Dungeon struct {
	Width, Height    int
	Link, Name, Desc string
	StartPosition    Position
	Tiles            [][]Tile
	Doors            []Door
	Keys             []Key
	Mobs             []Mob
}

// NewDungeon generates a new map initialized to the specified size
func NewDungeon(width, height int) *Dungeon {
	m := Dungeon{}

	m.Width = width
	m.Height = height
	m.Tiles = make([][]Tile, width)

	for i := range m.Tiles {
		m.Tiles[i] = make([]Tile, height)
	}

	return &m
}
