package file

import (
	"github.com/apsdsm/mapmaker/formats/placeholder"
)

func ValidateEntities(col *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {

	for i := 0; i < len(col.Mobs); i++ {
		mob := &col.Mobs[i]

		// check for prototype
		if mob.NeedsPrototype() && !col.HasMob(mob.Prototype) {
			e := Error{
				Message:    "mob '" + mob.Reference + "' requires prototype '" + mob.Prototype + "', which is not defined by any entity.",
				LineNumber: -1,
			}

			errors = append(errors, e)
		}

		for _, loot := range mob.ParsedLoot {
			if !col.HasItem(loot.Link) {
				e := Error{
					Message:    "mob '" + mob.Reference + "' requires item '" + loot.Link + "', which is not defined by any entity.",
					LineNumber: -1,
				}

				errors = append(errors, e)
			}
		}
	}

	for i := 0; i < len(col.Doors); i++ {
		door := &col.Doors[i]

		if door.NeedsKey() && !col.HasKey(door.Key) {
			e := Error{
				Message:    "door '" + door.Link + "' requires key '" + door.Key + "', which is not defined by any entity.",
				LineNumber: -1,
			}

			errors = append(errors, e)
		}
	}

	return errors, warnings
}
