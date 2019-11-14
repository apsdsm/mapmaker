package file

import "github.com/apsdsm/mapmaker/formats/placeholder"

// EntityValidator is responsible for making sure all entities are valid and can be used in a game.
type EntityValidator struct {
	Errors *ErrorList
}

// EntityValidatorConfig cotnains settings for EntityValidator.
type EntityValidatorConfig struct {
	Errors *ErrorList
}

// NewEntityValidator initializes and returns a new EntityValidator.
func NewEntityValidator(c EntityValidatorConfig) *EntityValidator {
	v := EntityValidator{
		Errors: c.Errors,
	}

	return &v
}

// Validate will check each entity in the list and make sure it contains no errors
func (v *EntityValidator) Validate(entities *placeholder.EntityList) error {

	// should validate something here...

	return nil
}
