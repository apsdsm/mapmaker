package file_test

import (
	"testing"

	"github.com/apsdsm/mapmaker/file"
	"github.com/stretchr/testify/assert"
)

func TestEntityImporter_ReadMobs(t *testing.T) {
	importer := file.NewEntityImporter(file.EntityImporterConfig{
		Errors: file.NewErrorList(),
	})

	err := importer.Read("../fixtures/entities/entities.yaml")
	assert.Nil(t, err)

	t.Run("it adds mob data", func(t *testing.T) {
		assert.Len(t, importer.Entities.Mobs, 1)

		mob1 := importer.Entities.Mob("goblin")
		assert.NotNil(t, mob1)

		// general statistics
		assert.Equal(t, "goblin_prot", mob1.Prototype)
		assert.Equal(t, 10, mob1.Hp)
		assert.Equal(t, 5, mob1.Mp)
		assert.Equal(t, "g", mob1.Rune)

		// loot
		assert.Len(t, mob1.Loot, 4)
		assert.Len(t, mob1.ParsedLoot, 4)

		assert.Equal(t, "copper_coin", mob1.ParsedLoot[0].Link)
		assert.Equal(t, 4, mob1.ParsedLoot[0].MinHeld)
		assert.Equal(t, 4, mob1.ParsedLoot[0].MaxHeld)

		assert.Equal(t, "silver_coin", mob1.ParsedLoot[1].Link)
		assert.Equal(t, 1, mob1.ParsedLoot[1].MinHeld)
		assert.Equal(t, 3, mob1.ParsedLoot[1].MaxHeld)

		assert.Equal(t, "gold_coin", mob1.ParsedLoot[2].Link)
		assert.Equal(t, 0, mob1.ParsedLoot[2].MinHeld)
		assert.Equal(t, 2, mob1.ParsedLoot[2].MaxHeld)

		assert.Equal(t, "axe", mob1.ParsedLoot[3].Link)
		assert.Equal(t, 1, mob1.ParsedLoot[3].MinHeld)
		assert.Equal(t, 1, mob1.ParsedLoot[3].MaxHeld)

		// events
		assert.Len(t, mob1.Events, 2)
		assert.Equal(t, "mob_on_die", mob1.Events["on_die"])
		assert.Equal(t, "mob_on_notice", mob1.Events["on_notice"])
	})
}

func TestEntityImporter_ReadDoors(t *testing.T) {

	importer := file.NewEntityImporter(file.EntityImporterConfig{
		Errors: file.NewErrorList(),
	})

	err := importer.Read("../fixtures/entities/entities.yaml")
	assert.Nil(t, err)

	t.Run("adds door data", func(t *testing.T) {
		door1 := importer.Entities.Door("locked_door")

		// general settings
		assert.Equal(t, "locked_door", door1.Reference)
		assert.Equal(t, true, door1.Locked)
		assert.Equal(t, "door_key", door1.Key)

		// events
		assert.Len(t, door1.Events, 1)
		assert.Equal(t, "door_on_try", door1.Events["on_try"])
	})

}

func TestEntityImporter_ReadItems(t *testing.T) {

	importer := file.NewEntityImporter(file.EntityImporterConfig{
		Errors: file.NewErrorList(),
	})

	err := importer.Read("../fixtures/entities/entities.yaml")
	assert.Nil(t, err)

	t.Run("loads item data", func(t *testing.T) {
		item1 := importer.Entities.Item("goboskull")

		assert.Len(t, importer.Entities.Items, 1)

		// data
		assert.Equal(t, "goboskull", item1.Reference)
		assert.Equal(t, "Gobo's Skull", item1.Name)
		assert.Equal(t, "It's Gobo's Skull.", item1.Desc)
		assert.Equal(t, "skull", item1.Type)
		assert.Equal(t, "1", item1.Size)
		assert.Equal(t, "true", item1.Uniq)
		// events
		assert.Len(t, item1.Events, 1)
		assert.Equal(t, "goboskull_on_use", item1.Events["on_use"])
	})
}
