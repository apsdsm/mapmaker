package file

import (
	"fmt"

	"github.com/apsdsm/mapmaker/formats/placeholder"
)

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

	// validate mobs
	for ref := range entities.Mobs {
		mob := entities.Mobs[ref]

		// check prototype
		if mob.NeedsPrototype() {
			if _, exists := entities.Mobs[mob.Prototype]; !exists {
				v.Errors.Add(Error{
					Message:   fmt.Sprintf("prototype not found: mobs.%s.prototype.%s", mob.Reference, mob.Prototype),
					IsWarning: false,
				})
			}
		}

		// check loot
		for _, loot := range mob.ParsedLoot {

			// make sure item exists
			if _, exists := entities.Items[loot.Link]; !exists {
				v.Errors.Add(Error{
					Message:   fmt.Sprintf("item not found: mobs.%s.loot.%s", mob.Reference, loot.Link),
					IsWarning: false,
				})
			}

			// make sure the monster has at least a chance of carrying the item
			if loot.MinHeld < 1 || loot.MaxHeld < loot.MinHeld {
				v.Errors.Add(Error{
					Message:   fmt.Sprintf("loot will never be carried: mobs.%s.loot.%s", mob.Reference, loot.Link),
					IsWarning: true,
				})
			}
		}
	}

	// validate doors
	for ref := range entities.Doors {
		door := entities.Doors[ref]

		if door.NeedsKey() {
			if _, exists := entities.Items[door.Key]; !exists {
				v.Errors.Add(Error{
					Message:   fmt.Sprintf("item not found: doors.%s.key.%s", door.Reference, door.Key),
					IsWarning: false,
				})
			}
		}
	}

	return nil
}
