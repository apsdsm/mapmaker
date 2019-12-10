package file_test

import (
	"testing"

	"github.com/apsdsm/mapmaker/file"
	"github.com/stretchr/testify/assert"
)

func TestDungeonImporter(t *testing.T) {
	importer := file.NewDungeonImporter(file.DungeonImporterConfig{})

	err := importer.Read("../fixtures/maps/meta_and_map.map")

	assert.NoError(t, err)

	t.Run("it loads a metadata", func(t *testing.T) {
		assert.Equal(t, "prison", importer.Dungeon.Link)
		assert.Equal(t, importer.Dungeon.Name, "The Jovian Prison")
		assert.Equal(t, importer.Dungeon.Description, "A gloomy building hidden deep in the Jovian woods.")
	})

	t.Run("it loads layout", func(t *testing.T) {
		assert.Equal(t, 11, importer.Dungeon.Width)
		assert.Equal(t, 7, importer.Dungeon.Height)
		assert.Equal(t, importer.Dungeon.Width, len(importer.Dungeon.Grid))

		for i := range importer.Dungeon.Grid {
			assert.Equal(t, importer.Dungeon.Height, len(importer.Dungeon.Grid[i]))
		}
	})

	t.Run("it loads line metadata", func(t *testing.T) {
		assert.Equal(t, "mob", importer.Dungeon.Grid[2][3].Type)
		assert.Equal(t, "mob1", importer.Dungeon.Grid[2][3].Link)

		assert.Equal(t, "door", importer.Dungeon.Grid[6][1].Type)
		assert.Equal(t, "door1", importer.Dungeon.Grid[6][1].Link)

		assert.Equal(t, "waypoint", importer.Dungeon.Grid[2][4].Type)
		assert.Equal(t, "waypoint1", importer.Dungeon.Grid[2][4].Link)

		assert.Equal(t, "item", importer.Dungeon.Grid[2][5].Type)
		assert.Equal(t, "item1", importer.Dungeon.Grid[2][5].Link)
	})

	t.Run("it assigns a start position", func(t *testing.T) {
		assert.Equal(t, 2, importer.Dungeon.StartPosition.X)
		assert.Equal(t, 1, importer.Dungeon.StartPosition.Y)
	})

	t.Run("it sets rule value to ' ' if cell contains entity", func(t *testing.T) {
		assert.Equal(t, ' ', importer.Dungeon.Grid[2][3].Rune, "mob entity rune should be blank")
		assert.Equal(t, ' ', importer.Dungeon.Grid[6][1].Rune, "door entity rune should be blank")
		assert.Equal(t, ' ', importer.Dungeon.Grid[2][4].Rune, "waypoint entity rune should be blank")
		assert.Equal(t, ' ', importer.Dungeon.Grid[2][5].Rune, "item entity rune should be blank")
	})
}
