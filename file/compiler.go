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
	"github.com/apsdsm/mapmaker/placeholders"
)

// CompileLevel converts placeholder information to a map file that can be saved as json
func CompileLevel(metaData *placeholders.Meta, mapData *placeholders.Map, entities *placeholders.EntityCollection) *placeholders.Dungeon {
	m := placeholders.NewDungeon(mapData.Width, mapData.Height)

	// copy tile data
	// - set floor tiles as walkable
	// - set spawn data for tiles
	for x := 0; x < mapData.Width; x++ {
		for y := 0; y < mapData.Height; y++ {
			m.Tiles[x][y].Rune = mapData.Grid[x][y].Rune

			if mapData.Grid[x][y].Rune == ' ' {
				m.Tiles[x][y].Walkable = true
			}
			if mapData.Grid[x][y].Type == "mob" {
				m.Tiles[x][y].Spawn = mapData.Grid[x][y].Link
			}
		}
	}

	// start position
	if mapData.StartPosition != nil {
		m.StartPosition = *mapData.StartPosition
	}

	// copy meta
	m.Link = metaData.Link
	m.Name = metaData.Name
	m.Desc = metaData.Desc

	// copy entities
	m.Doors = entities.Doors
	m.Mobs = entities.Mobs
	m.Keys = entities.Keys

	return m
}
