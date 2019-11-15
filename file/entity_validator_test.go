package file_test

import (
	"testing"

	"github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/formats/placeholder"
	"github.com/stretchr/testify/assert"
)

func TestValidator_MobPrototypeChecking(t *testing.T) {
	validator := file.NewEntityValidator(file.EntityValidatorConfig{
		Errors: file.NewErrorList(),
	})

	entities := placeholder.NewEntityList()

	// mob with no prototype
	entities.Mobs["mob1"] = &placeholder.Mob{
		Name:      "No Prototype",
		Reference: "mob1",
		Prototype: "does-not-exist",
		Rune:      "g",
	}

	err := validator.Validate(entities)
	assert.NoError(t, err)

	assertErrorExists(t, validator.Errors, "prototype not found: mobs.mob1.prototype.does-not-exist", false)
}

func TestValidator_MobItemChecking(t *testing.T) {
	validator := file.NewEntityValidator(file.EntityValidatorConfig{
		Errors: file.NewErrorList(),
	})

	entities := placeholder.NewEntityList()

	// mob with missing loot
	entities.Mobs["mob1"] = &placeholder.Mob{
		Name:      "Mising Items",
		Reference: "mob1",
		Rune:      "g",
		ParsedLoot: []placeholder.Loot{
			{
				Link: "does-not-exist",
			},
		},
	}

	err := validator.Validate(entities)
	assert.NoError(t, err)

	assertErrorExists(t, validator.Errors, "item not found: mobs.mob1.loot.does-not-exist", false)
	assertErrorExists(t, validator.Errors, "loot will never be carried: mobs.mob1.loot.does-not-exist", true)
}

func TestValidator_DoorChecking(t *testing.T) {
	validator := file.NewEntityValidator(file.EntityValidatorConfig{
		Errors: file.NewErrorList(),
	})

	entities := placeholder.NewEntityList()

	// mob with missing loot
	entities.Doors["door1"] = &placeholder.Door{
		Reference: "door1",
		Locked:    false,
		Key:       "does-not-exist",
	}

	err := validator.Validate(entities)
	assert.NoError(t, err)

	assertErrorExists(t, validator.Errors, "item not found: doors.door1.key.does-not-exist", false)
}

func assertErrorExists(t *testing.T, errors *file.ErrorList, message string, warning bool) {
	all := errors.All()

	for e := range all {
		if all[e].Message == message && all[e].IsWarning == warning {
			return
		}
	}

	assert.Fail(t, "did not find error: "+message)
}
