package file_test

import (
	"testing"

	"github.com/apsdsm/mapmaker/file"
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
		// [#][ ]
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

	// fix tests and code for tests
	t.Run("it adds entities to the map", func(t *testing.T) {

		compiler := file.NewCompiler(file.CompilerConfig{})

		// dungeon layout:
		// [d][m][i]
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
	})

	//
	//It("adds mob loot", func() {
	//	dungeon := placeholder.NewMap(0, 0)
	//	meta := placeholder.Meta{}
	//
	//	mob := placeholder.Mob{
	//		ParsedLoot: []placeholder.Loot{
	//			placeholder.Loot{
	//				Link:    "loot",
	//				MinHeld: 1,
	//				MaxHeld: 2,
	//			},
	//		},
	//	}
	//
	//	loot := placeholder.Item{
	//		//Link: "loot",
	//	}
	//
	//	entities := placeholder.NewEntityCollection()
	//	entities.AddMobs(mob)
	//	entities.AddItems(loot)
	//
	//	compiled := Compile(&meta, dungeon, &entities)
	//
	//	Expect(len(compiled.Mobs)).To(Equal(1))
	//	Expect(len(compiled.Items)).To(Equal(1))
	//
	//	Expect(len(compiled.Mobs[0].Loot)).To(Equal(1))
	//	Expect(compiled.Mobs[0].Loot[0].Link).To(Equal("loot"))
	//	Expect(compiled.Mobs[0].Loot[0].MinHeld).To(Equal(1))
	//	Expect(compiled.Mobs[0].Loot[0].MinHeld).To(Equal(2))
	//
	//})
	//
	//It("sets the spawn values for each cell", func() {
	//	dungeon := placeholder.NewMap(1, 1)
	//	dungeon.Grid[0][0] = &placeholder.Cell{
	//		Type: "mob",
	//		Link: "mob_link",
	//	}
	//
	//	meta := placeholder.Meta{}
	//
	//	entities := placeholder.NewEntityCollection()
	//
	//	compiled := Compile(&meta, dungeon, &entities)
	//
	//	Expect(compiled.Tiles[0][0].Spawn).To(Equal("mob_link"))
	//})
	//
	//It("Adds the start position", func() {
	//	dungeon := placeholder.NewMap(0, 0)
	//
	//	dungeon.StartPosition = &placeholder.Position{1, 1}
	//
	//	meta := placeholder.Meta{}
	//
	//	entities := placeholder.NewEntityCollection()
	//
	//	compiled := Compile(&meta, dungeon, &entities)
	//
	//	Expect(compiled.StartPosition).To(Equal(output.Position{1, 1}))
	//})
}
