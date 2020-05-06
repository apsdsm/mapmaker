package file_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/apsdsm/mapmaker/formats/placeholder"

	"github.com/apsdsm/mapmaker/file"
)

func TestDungeonValidator_Validate(t *testing.T) {

	t.Run("it returns error if cell contains mob that doesn't exist", func(t *testing.T) {

		dungeon := placeholder.NewDungeon(1, 1)
		dungeon.Name = "test"
		dungeon.Grid[0][0] = &placeholder.Cell{
			Type:      "mob",
			Link:      "does-not-exist",
			Annotated: true,
			X:         0,
			Y:         0,
		}

		dungeon.StartPosition = &placeholder.Position{0, 0}

		entities := placeholder.NewEntityList()

		errors := file.NewErrorList()

		validator := file.NewDungeonValidator(file.DungeonValidatorConfig{
			Errors: errors,
		})

		err := validator.Validate(dungeon, entities)

		assert.NoError(t, err)
		assert.Len(t, errors.All(), 1)
		assert.Equal(t, 1, errors.All()[0].LineNumber)
		assert.Equal(t, "mob does-not-exist is not defined by any entity", errors.All()[0].Message)
		assert.False(t, errors.All()[0].IsWarning)
	})

	t.Run("it returns error if cell contains door that doesn't exist", func(t *testing.T) {
		dungeon := placeholder.NewDungeon(1, 1)
		dungeon.Name = "test"
		dungeon.Grid[0][0] = &placeholder.Cell{
			Type:      "door",
			Link:      "does-not-exist",
			Annotated: true,
			X:         0,
			Y:         0,
		}

		dungeon.StartPosition = &placeholder.Position{0, 0}

		entities := placeholder.NewEntityList()
		errors := file.NewErrorList()

		validator := file.NewDungeonValidator(file.DungeonValidatorConfig{
			Errors: errors,
		})

		err := validator.Validate(dungeon, entities)

		assert.NoError(t, err)
		assert.Len(t, errors.All(), 1)
		assert.Equal(t, 1, errors.All()[0].LineNumber)
		assert.Equal(t, "door does-not-exist is not defined by any entity", errors.All()[0].Message)
		assert.False(t, errors.All()[0].IsWarning)
	})

	t.Run("it returns an error if there is no start position", func(t *testing.T) {
		dungeon := placeholder.NewDungeon(1, 1)
		dungeon.Name = "test"
		dungeon.Grid[0][0] = &placeholder.Cell{
			X: 0,
			Y: 0,
		}

		entities := placeholder.NewEntityList()
		errors := file.NewErrorList()

		validator := file.NewDungeonValidator(file.DungeonValidatorConfig{
			Errors: errors,
		})

		err := validator.Validate(dungeon, entities)

		assert.NoError(t, err)
		assert.Len(t, errors.All(), 1)
		assert.Equal(t, -1, errors.All()[0].LineNumber)
		assert.Equal(t, "map has no start position", errors.All()[0].Message)
		assert.False(t, errors.All()[0].IsWarning)
	})

	t.Run("it returns a warning if there is no name", func(t *testing.T) {
		dungeon := placeholder.NewDungeon(1, 1)
		dungeon.StartPosition = &placeholder.Position{0, 0}
		dungeon.Grid[0][0] = &placeholder.Cell{
			X: 0,
			Y: 0,
		}

		entities := placeholder.NewEntityList()
		errors := file.NewErrorList()

		validator := file.NewDungeonValidator(file.DungeonValidatorConfig{
			Errors: errors,
		})

		err := validator.Validate(dungeon, entities)

		assert.NoError(t, err)
		assert.Len(t, errors.All(), 1)
		assert.Equal(t, -1, errors.All()[0].LineNumber)
		assert.Equal(t, "map has no name", errors.All()[0].Message)
		assert.True(t, errors.All()[0].IsWarning)
	})
}
