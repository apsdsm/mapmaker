package file_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/apsdsm/mapmaker/formats/placeholder"

	"github.com/apsdsm/mapmaker/file"
)

func TestDungeonValidator_Validate(t *testing.T) {

	//var (
	//	errors []file.Error
	//	//	warnings []file.Error
	//)

	//BeforeEach(func() {
	//
	//	// empty errors/warnings
	//	errors = make([]file.Error, 0, 0)
	//	//	warnings = make([]file.Error, 0, 0)
	//})
	//

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
		assert.Equal(t, "mob does-not-exist is not defined by any entity.", errors.All()[0].Message)

		//Expect(len(errors)).To(Equal(1))
		//Expect(errors[0].LineNumber).To(Equal(1))
		//Expect(errors[0].Message).To(Equal("mob is not defined by any entity."))
	})

	//
	//It("returns error if cell contains door that doesn't exist", func() {
	//	meta := placeholder.Meta{}
	//	meta.Name = "foo dungeon"
	//
	//	dungeon := placeholder.NewMap(1, 1)
	//	dungeon.Grid[0][0] = placeholder.EmptyCell()
	//	dungeon.Grid[0][0].Annotated = true
	//	dungeon.Grid[0][0].Link = "door_link"
	//	dungeon.Grid[0][0].Type = "door"
	//	dungeon.StartPosition = &placeholder.Position{0, 0}
	//
	//	// empty entity collection
	//	entities := placeholder.NewEntityCollection()
	//
	//	errors, _ = file.ValidateDungeon(&meta, dungeon, &entities)
	//
	//	Expect(len(errors)).To(Equal(1))
	//	Expect(errors[0].LineNumber).To(Equal(1))
	//	Expect(errors[0].Message).To(Equal("door_link is not defined by any entity."))
	//})
	//
	//It("returns error if there is no start position", func() {
	//	meta := placeholder.Meta{}
	//	meta.Name = "foo dungeon"
	//
	//	// dungeon with no start position
	//	dungeon := placeholder.NewMap(1, 1)
	//	dungeon.Grid[0][0] = placeholder.EmptyCell()
	//
	//	entities := placeholder.NewEntityCollection()
	//
	//	errors, _ := file.ValidateDungeon(&meta, dungeon, &entities)
	//
	//	Expect(len(errors)).To(Equal(1))
	//	Expect(errors[0].LineNumber).To(Equal(-1))
	//	Expect(errors[0].Message).To(Equal("map has no start position."))
	//})
	//
	//It("returns warning if there is no name", func() {
	//	meta := placeholder.Meta{}
	//
	//	dungeon := placeholder.NewMap(1, 1)
	//	dungeon.Grid[0][0] = placeholder.EmptyCell()
	//	dungeon.StartPosition = &placeholder.Position{0, 0}
	//
	//	entities := placeholder.NewEntityCollection()
	//
	//	errors, _ := file.ValidateDungeon(&meta, dungeon, &entities)
	//
	//	Expect(len(errors)).To(Equal(1))
	//	//Expect(errors[0].IsWarning).To(BeTrue())
	//	Expect(errors[0].LineNumber).To(Equal(-1))
	//	Expect(errors[0].Message).To(Equal("map has no name."))
	//})

}
