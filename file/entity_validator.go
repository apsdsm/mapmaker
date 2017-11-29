package file

import (
	"github.com/apsdsm/mapmaker/formats/placeholder"
)

func ValidateEntities(col *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {

	for i := 0; i < len(col.Mobs); i++ {

		mob := &col.Mobs[i]

		// check for prototype
		if mob.NeedsPrototype() && !col.HasMob(mob.Prot) {
			e := Error{
				Message:    "mob '" + mob.Link + "' requires prototype '" + mob.Prot + "', which is not defined by any entity.",
				LineNumber: -1,
			}

			errors = append(errors, e)
		}

		for _, loot := range mob.ParsedLoot {
			if !col.HasItem(loot.Link) {
				e := Error{
					Message:    "mob '" + mob.Link + "' requires item '" + loot.Link + "', which is not defined by any entity.",
					LineNumber: -1,
				}

				errors = append(errors, e)
			}
		}
	}

	return errors, warnings
}
