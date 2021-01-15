package file_test

import (
	"testing"

	"github.com/apsdsm/mapmaker/file"
	output_format "github.com/apsdsm/mapmaker/formats/output"
	"github.com/apsdsm/mapmaker/formats/placeholder"
	"github.com/stretchr/testify/assert"
)

func TestCompile(t *testing.T) {

	t.Run("it adds metadata to the map", func(t *testing.T) {
		compiler := file.NewCompiler(file.CompilerConfig{})

		dungeon := placeholder.NewDungeon(0, 0)
		dungeon.Link = "meta_link"
		dungeon.Name = "meta_name"
		dungeon.Description = "meta_desc"

		entities := &placeholder.EntityCollection{}

		output, err := compiler.Compile(dungeon, entities)

		assert.NoError(t, err)
		assert.Equal(t, "meta_link", output.Link)
		assert.Equal(t, "meta_name", output.Name)
		assert.Equal(t, "meta_desc", output.Desc)
	})

	t.Run("it converts runes into walls and floors", func(t *testing.T) {
		compiler := file.NewCompiler(file.CompilerConfig{})

		// dungeon layout:
		// [#]
		// [ ]
		//
		dungeon := placeholder.NewDungeon(2, 1)
		dungeon.Grid[0][0] = &placeholder.Cell{Rune: '#'}
		dungeon.Grid[0][1] = &placeholder.Cell{Rune: ' '}

		entities := &placeholder.EntityCollection{}

		output, err := compiler.Compile(dungeon, entities)

		assert.NoError(t, err)
		assert.False(t, output.Tiles[0][0].Walkable)
		assert.True(t, output.Tiles[0][1].Walkable)
	})

	t.Run("it adds dungeon positional data", func(t *testing.T) {
		compiler := file.NewCompiler(file.CompilerConfig{})

		// dungeon layout:
		// [#]
		// [ ]
		//
		dungeon := placeholder.NewDungeon(2, 1)
		dungeon.Grid[0][0] = &placeholder.Cell{Rune: '#'}
		dungeon.Grid[0][1] = &placeholder.Cell{Rune: ' '}
		dungeon.StartPosition = &placeholder.Position{X: 0, Y: 1}

		entities := &placeholder.EntityCollection{}

		output, err := compiler.Compile(dungeon, entities)

		assert.NoError(t, err)
		assert.Equal(t, output_format.Position{X: 0, Y: 1}, output.StartPosition)
	})

	// fix tests and code for tests
	t.Run("it adds entities to the map", func(t *testing.T) {
		compiler := file.NewCompiler(file.CompilerConfig{})

		// dungeon layout:
		// [d]
		// [m]
		// [i]
		//
		dungeon := placeholder.NewDungeon(3, 1)

		dungeon.Grid[0][0] = &placeholder.Cell{
			Rune: ' ',
			Type: "door",
			Link: "door1",
		}

		dungeon.Grid[0][1] = &placeholder.Cell{
			Rune: ' ',
			Type: "mob",
			Link: "mob1",
		}

		dungeon.Grid[0][2] = &placeholder.Cell{
			Rune: ' ',
			Type: "item",
			Link: "item1",
		}

		entities := &placeholder.EntityCollection{}

		entities.AddDoors(placeholder.Door{
			Reference: "door1",
		})

		entities.AddMobs(placeholder.Mob{
			Reference: "mob1",
			ParsedLoot: []placeholder.Loot{
				{
					Link:    "item1",
					MaxHeld: 2,
					MinHeld: 1,
				},
			},
		})

		entities.AddItems(placeholder.Item{
			Reference: "item1",
		})

		output, err := compiler.Compile(dungeon, entities)

		assert.NoError(t, err)
		assert.Equal(t, "door1", output.Tiles[0][0].Spawn)
		assert.True(t, output.Tiles[0][0].Walkable)

		assert.Equal(t, "mob1", output.Tiles[0][1].Spawn)
		assert.True(t, output.Tiles[0][1].Walkable)

		assert.Equal(t, "item1", output.Tiles[0][2].Spawn)
		assert.True(t, output.Tiles[0][2].Walkable)

		assert.Len(t, output.Mobs[0].Loot, 1)
		assert.Equal(t, "item1", output.Mobs[0].Loot[0].Link)
		assert.Equal(t, 1, output.Mobs[0].Loot[0].MinHeld)
		assert.Equal(t, 2, output.Mobs[0].Loot[0].MaxHeld)
	})
}
