package file_test

import (
	"github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/formats/placeholder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EntityValidator", func() {

	var (
		errors   []file.Error
		warnings []file.Error
	)

	BeforeEach(func() {
		errors = make([]file.Error, 0, 0)
		warnings = make([]file.Error, 0, 0)
	})

	It("returns an error if a mob references a prototype that doesn't exist", func() {

		// mob with missing prototype
		mob := placeholder.Mob{}
		mob.Link = "mob"
		mob.Prot = "foo"

		// add mob to entities
		entities := placeholder.NewEntityCollection()
		entities.AddMobs(mob)

		errors, _ = file.ValidateEntities(&entities, errors, warnings)

		Expect(len(errors)).To(Equal(1))
		Expect(errors[0].LineNumber).To(Equal(-1))
		Expect(errors[0].Message).To(Equal("mob 'mob' requires prototype 'foo', which is not defined by any entity."))
	})

	It("returns an error if mob references item that doesn't exist", func() {

		// mob with missing item
		mob := placeholder.Mob{}
		mob.Link = "mob"
		mob.ParsedLoot = make([]placeholder.Loot, 1)
		mob.ParsedLoot[0].Link = "foo"
		mob.ParsedLoot[0].MaxHeld = 1
		mob.ParsedLoot[0].MinHeld = 1

		entities := placeholder.NewEntityCollection()
		entities.AddMobs(mob)

		errors, _ = file.ValidateEntities(&entities, errors, warnings)

		Expect(len(errors)).To(Equal(1))
		Expect(errors[0].LineNumber).To(Equal(-1))
		Expect(errors[0].Message).To(Equal("mob 'mob' requires item 'foo', which is not defined by any entity."))
	})

	It("returns error if door uses a key that doesn't exist", func() {

		// door with missing key
		door := placeholder.Door{}
		door.Link = "door_link"
		door.Key = "key_link"

		entities := placeholder.NewEntityCollection()
		entities.AddDoors(door)

		errors, _ := file.ValidateEntities(&entities, errors, warnings)

		Expect(len(errors)).To(Equal(1))
		Expect(errors[0].LineNumber).To(Equal(-1))
		Expect(errors[0].Message).To(Equal("door 'door_link' requires key 'key_link', which is not defined by any entity."))
	})
})
