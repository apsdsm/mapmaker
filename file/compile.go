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
	"github.com/apsdsm/mapmaker/formats/output"
	"github.com/apsdsm/mapmaker/formats/placeholder"
)

// Compile converts placeholder information to a map file that can be saved as json
func Compile(metaData *placeholder.Meta, mapData *placeholder.Map, entities *placeholder.EntityCollection) *output.Dungeon {
	m := output.NewDungeon(mapData.Width, mapData.Height)

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
	m.Mobs = mobsToJSON(entities.Mobs)
	m.Keys = keysToJSON(entities.Keys)
	m.Items = itemsToJSON(entities.Items)

	return m
}

func positionToJson(position placeholder.Position) output.Position {
	return output.Position{
		X: position.X,
		Y: position.Y,
	}
}

func doorsToJson(doors []placeholder.Door) []output.Door {
	d := make([]output.Door, len(doors))

	for i := range doors {
		d[i] = output.Door{
			Link:   doors[i].Link,
			Locked: doors[i].Locked,
			Key:    doors[i].Key,
			OnTry:  doors[i].OnTry,
		}
	}

	return d
}

func mobsToJSON(mobs []placeholder.Mob) []output.Mob {
	m := make([]output.Mob, len(mobs))

	for i := range mobs {

		newMob := output.Mob{
			Name: mobs[i].Name,
			Link: mobs[i].Reference,
			Prot: mobs[i].Prototype,
			Rune: mobs[i].Rune,
			Loot: make([]output.Loot, len(mobs[i].ParsedLoot)),
		}

		for j := range mobs[i].ParsedLoot {
			newMob.Loot[j] = output.Loot{
				Link:    mobs[i].ParsedLoot[j].Link,
				MinHeld: mobs[i].ParsedLoot[j].MinHeld,
				MaxHeld: mobs[i].ParsedLoot[j].MaxHeld,
			}
		}

		m[i] = newMob
	}

	return m
}

func keysToJSON(keys []placeholder.Key) []output.Key {
	k := make([]output.Key, len(keys))

	for i := range keys {
		k[i] = output.Key{
			Name: keys[i].Name,
			Link: keys[i].Link,
			Desc: keys[i].Desc,
		}
	}

	return k
}

func itemsToJSON(items []placeholder.Item) []output.Item {
	r := make([]output.Item, len(items))

	for i := range items {
		r[i] = output.Item{
			Link: items[i].Link,
			Type: items[i].Type,
			Name: items[i].Name,
			Desc: items[i].Desc,
			Size: items[i].Size,
			Uniq: items[i].Uniq,
		}
	}

	return r
}
