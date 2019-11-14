package file_test

import (
	"testing"

	"github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/formats/placeholder"
	"github.com/stretchr/testify/assert"
)

func TestValidator_PrototypeChecking(t *testing.T) {
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

	assertErrorExists(t, validator.Errors, "prototype not found: does-not-exist ==> mob1", false)
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
