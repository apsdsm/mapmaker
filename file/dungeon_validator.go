package file

import (
	"fmt"

	"github.com/apsdsm/mapmaker/formats/placeholder"
)

type DungeonValidator struct {
	Errors *ErrorList
}

type DungeonValidatorConfig struct {
	Errors *ErrorList
}

func NewDungeonValidator(c DungeonValidatorConfig) *DungeonValidator {
	v := DungeonValidator{
		Errors: c.Errors,
	}

	return &v
}

func (v *DungeonValidator) Validate(dungeon *placeholder.Dungeon, entities *placeholder.EntityList) error {

	// check for name
	if dungeon.Name == "" {
		v.Errors.Add(Error{
			Message:    "map has no name.",
			LineNumber: -1,
			IsWarning:  true,
		})
	}

	//// check for start point
	//if levelData.StartPosition == nil {
	//	err := Error{
	//		Message:    "map has no start position.",
	//		LineNumber: -1,
	//	}
	//	errors = append(errors, err)
	//}
	//

	// validate each cell
	for x := 0; x < dungeon.Width; x++ {
		for y := 0; y < dungeon.Height; y++ {
			err := v.validateCell(dungeon.Grid[x][y], y+1, entities)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (v *DungeonValidator) validateCell(cell *placeholder.Cell, line int, entities *placeholder.EntityList) error {
	if !cell.Annotated {
		return nil
	}

	if cell.Type == "mob" {
		if !entities.HasMob(cell.Link) {
			v.Errors.Add(Error{
				Message:    fmt.Sprintf("mob %s is not defined by any entity.", cell.Link),
				LineNumber: line,
			})
		}
	}

	if cell.Type == "door" {
		if !entities.HasDoor(cell.Link) {
			v.Errors.Add(Error{
				Message:    fmt.Sprintf("door %s is not defined by any entity.", cell.Link),
				LineNumber: line,
			})
		}
	}

	return nil
}
