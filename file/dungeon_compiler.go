package file

import (
	"github.com/apsdsm/mapmaker/formats/output"
	"github.com/apsdsm/mapmaker/formats/placeholder"
)

// A Compiler takes a validated placeholder.Dungeon and placeholder.EntityList and compiles them into a single .dng file.
type Compiler struct {
}

// CompilerConfig contains the config data for a Compiler
type CompilerConfig struct {
}

// NewCompiler will create and initialize a new Compiler
func NewCompiler(config CompilerConfig) *Compiler {
	c := Compiler{}
	return &c
}

// Compile will compile the raw map data into a compiled map, but what it should do is compile an array of dungeons together,
// with the entities, as separate files!
func (c *Compiler) Compile(dungeon *placeholder.Dungeon, entities *placeholder.EntityCollection) (*output.Dungeon, error) {
	m := output.NewDungeon(dungeon.Width, dungeon.Height)

	// copy tile data
	// - set floor tiles as walkable
	// - set spawn data for tiles
	for x := 0; x < dungeon.Width; x++ {
		for y := 0; y < dungeon.Height; y++ {
			m.Tiles[x][y].Rune = dungeon.Grid[x][y].Rune

			if dungeon.Grid[x][y].Rune == ' ' {
				m.Tiles[x][y].Walkable = true
			}

			switch dungeon.Grid[x][y].Type {
			case "mob", "door", "item":
				m.Tiles[x][y].Spawn = dungeon.Grid[x][y].Link
				m.Tiles[x][y].Walkable = true
			}
		}
	}

	// start position
	if dungeon.StartPosition != nil {
		m.StartPosition = positionToJson(*dungeon.StartPosition)
	}

	// copy meta
	m.Link = dungeon.Link
	m.Name = dungeon.Name
	m.Desc = dungeon.Description

	// copy entities
	m.Doors = compileDoors(entities.Doors)
	m.Mobs = compilesMobs(entities.Mobs)
	m.Keys = compileKeys(entities.Keys)
	m.Items = compileItems(entities.Items)

	return m, nil
}

func positionToJson(position placeholder.Position) output.Position {
	return output.Position{
		X: position.X,
		Y: position.Y,
	}
}

func compileDoors(doors []placeholder.Door) []output.Door {
	d := make([]output.Door, len(doors))

	for i := range doors {
		d[i] = output.Door{
			// Link:   doors[i].Link,
			Locked: doors[i].Locked,
			Key:    doors[i].Key,
			// OnTry:  doors[i].OnTry,
		}
	}

	return d
}

func compilesMobs(mobs []placeholder.Mob) []output.Mob {
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

func compileKeys(keys []placeholder.Key) []output.Key {
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

func compileItems(items []placeholder.Item) []output.Item {
	r := make([]output.Item, len(items))

	for i := range items {
		r[i] = output.Item{
			// Link: items[i].Link,
			Type: items[i].Type,
			Name: items[i].Name,
			Desc: items[i].Desc,
			Size: items[i].Size,
			Uniq: items[i].Uniq,
		}
	}

	return r
}
