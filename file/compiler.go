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
	"github.com/apsdsm/mapmaker/formats/json_format"
	"github.com/apsdsm/mapmaker/formats/placeholder"
)

// CompileLevel converts placeholder information to a map file that can be saved as json
func CompileLevel(metaData *placeholder.Meta, mapData *placeholder.Map, entities *placeholder.EntityCollection) *json_format.Dungeon {
	m := json_format.NewDungeon(mapData.Width, mapData.Height)

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
		m.StartPosition = positionToJson(*mapData.StartPosition)
	}

	// copy meta
	m.Link = metaData.Link
	m.Name = metaData.Name
	m.Desc = metaData.Desc

	// copy entities
	m.Doors = doorsToJson(entities.Doors)
	m.Mobs = mobsToJson(entities.Mobs)
	m.Keys = keysToJson(entities.Keys)
	m.Items = itemsToJson(entities.Items)

	return m
}

func positionToJson(position placeholder.Position) json_format.Position {
	return json_format.Position{
		X: position.X,
		Y: position.Y,
	}
}

func doorsToJson(doors []placeholder.Door) []json_format.Door {
	d := make([]json_format.Door, len(doors))

	for i := range doors {
		d[i] = json_format.Door{
			Link:   doors[i].Link,
			Locked: doors[i].Locked,
			Key:    doors[i].Key,
			OnTry:  doors[i].OnTry,
		}
	}

	return d
}

func mobsToJson(mobs []placeholder.Mob) []json_format.Mob {
	m := make([]json_format.Mob, len(mobs))

	for i := range mobs {
		m[i] = json_format.Mob{
			Name: mobs[i].Name,
			Link: mobs[i].Link,
			Prot: mobs[i].Prot,
			Rune: mobs[i].Rune,
		}
	}

	return m
}

func keysToJson(keys []placeholder.Key) []json_format.Key {
	k := make([]json_format.Key, len(keys))

	for i := range keys {
		k[i] = json_format.Key{
			Name: keys[i].Name,
			Link: keys[i].Link,
			Desc: keys[i].Desc,
		}
	}

	return k
}

func itemsToJson(items []placeholder.Item) []json_format.Item {
	r := make([]json_format.Item, len(items))

	for i := range items {
		r[i] = json_format.Item{
			Link: r[i].Link,
			Type: r[i].Type,
			Name: r[i].Name,
			Desc: r[i].Desc,
			Size: r[i].Size,
			Uniq: r[i].Uniq,
		}
	}

	return r
}
