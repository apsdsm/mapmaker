// Copyright 2017 Nick del Pozo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/formats/placeholder"
)

var _ = Describe("EntityImporter", func() {

	var (
		entities placeholder.EntityCollection
		errors   []file.Error
		warnings []file.Error
	)

	BeforeEach(func() {
		entities = placeholder.NewEntityCollection()
		errors = make([]file.Error, 0, 0)
		warnings = make([]file.Error, 0, 0)
	})

	Context("importing a valid mob", func() {
		BeforeEach(func() {
			errors, warnings = file.AddEntityToCollection("../fixtures/entities/mob.mob.yaml", &entities, errors, warnings)
		})

		It("adds mob data", func() {
			Expect(len(entities.Mobs)).To(Equal(1))
			Expect(entities.Mobs[0].Name).To(Equal("Mob Name"))
			Expect(entities.Mobs[0].Link).To(Equal("mob_link"))
			Expect(entities.Mobs[0].Prot).To(Equal("mob_prot"))
			Expect(entities.Mobs[0].Hp).To(Equal("1"))
			Expect(entities.Mobs[0].Mp).To(Equal("2"))
			Expect(len(entities.Mobs[0].Loot)).To(Equal(3))
		})

		It("adds mob loot", func() {
			Expect(len(entities.Mobs[0].ParsedLoot)).To(Equal(3))
			Expect(entities.Mobs[0].ParsedLoot[0].Link).To(Equal("copper"))
			Expect(entities.Mobs[0].ParsedLoot[0].MinHeld).To(Equal(10))
			Expect(entities.Mobs[0].ParsedLoot[0].MaxHeld).To(Equal(10))

			Expect(entities.Mobs[0].ParsedLoot[1].Link).To(Equal("gold"))
			Expect(entities.Mobs[0].ParsedLoot[1].MinHeld).To(Equal(1))
			Expect(entities.Mobs[0].ParsedLoot[1].MaxHeld).To(Equal(2))

			Expect(entities.Mobs[0].ParsedLoot[2].Link).To(Equal("axe"))
			Expect(entities.Mobs[0].ParsedLoot[2].MinHeld).To(Equal(1))
			Expect(entities.Mobs[0].ParsedLoot[2].MaxHeld).To(Equal(1))
		})

		It("adds no errors", func() {
			Expect(len(errors)).To(BeZero())
		})

		It("adds no warnings", func() {
			Expect(len(warnings)).To(BeZero())
		})
	})

	Context("importing mob with broken item entry", func() {
		BeforeEach(func() {
			errors, warnings = file.AddEntityToCollection("../fixtures/entities/broken.mob.yaml", &entities, errors, warnings)
		})

		It("adds item errors", func() {
			Expect(len(errors)).To(Equal(3))
			Expect(errors[0].LineNumber).To(Equal(-1))
			Expect(errors[0].FileName).To(Equal("../fixtures/entities/broken.mob.yaml"))
			Expect(errors[0].Message).To(Equal("invalid item syntax: 'error1[10~'. Use notation: 'object_link[count]', 'object_link[min:max]', or 'object_link'"))

			Expect(errors[1].LineNumber).To(Equal(-1))
			Expect(errors[1].FileName).To(Equal("../fixtures/entities/broken.mob.yaml"))
			Expect(errors[1].Message).To(Equal("invalid item syntax: 'error2[10'. Use notation: 'object_link[count]', 'object_link[min:max]', or 'object_link'"))

			Expect(errors[2].LineNumber).To(Equal(-1))
			Expect(errors[2].FileName).To(Equal("../fixtures/entities/broken.mob.yaml"))
			Expect(errors[2].Message).To(Equal("invalid item syntax: 'error3['. Use notation: 'object_link[count]', 'object_link[min:max]', or 'object_link'"))
		})
	})

	Context("imorting valid door", func() {
		It("loads a yaml file of door", func() {
			source := "../fixtures/entities/door.door.yaml"

			entities := placeholder.NewEntityCollection()
			errors := make([]file.Error, 0, 0)
			warnings := make([]file.Error, 0, 0)
			errors, warnings = file.AddEntityToCollection(source, &entities, errors, warnings)

			Expect(len(entities.Doors)).To(Equal(1))
			Expect(entities.Doors[0].Link).To(Equal("door_link"))
			Expect(entities.Doors[0].Locked).To(Equal(true))
			Expect(entities.Doors[0].Key).To(Equal("door_key"))
			Expect(entities.Doors[0].OnTry).To(Equal("door_ontry"))
		})
	})

	Context("importing valid key", func() {
		It("loads a yaml file of key", func() {
			source := "../fixtures/entities/key.key.yaml"

			entities := placeholder.NewEntityCollection()
			errors := make([]file.Error, 0, 0)
			warnings := make([]file.Error, 0, 0)
			errors, warnings = file.AddEntityToCollection(source, &entities, errors, warnings)

			Expect(len(entities.Keys)).To(Equal(1))
			Expect(entities.Keys[0].Name).To(Equal("Normal Key"))
			Expect(entities.Keys[0].Link).To(Equal("normal_key"))
			Expect(entities.Keys[0].Desc).To(Equal("A normal key"))
		})
	})

	Context("importing valid item", func() {
		It("loads yaml for item", func() {
			source := "../fixtures/entities/key.item.yaml"

			entities := placeholder.NewEntityCollection()
			errors := make([]file.Error, 0, 0)
			warnings := make([]file.Error, 0, 0)
			errors, warnings = file.AddEntityToCollection(source, &entities, errors, warnings)

			Expect(len(entities.Items)).To(Equal(1))
			Expect(entities.Items[0].Name).To(Equal("A Small Key"))
			Expect(entities.Items[0].Link).To(Equal("small_key"))
			Expect(entities.Items[0].Desc).To(Equal("A regular key"))
			Expect(entities.Items[0].Type).To(Equal("key"))
			Expect(entities.Items[0].Size).To(Equal("1"))
			Expect(entities.Items[0].Uniq).To(Equal("true"))
		})
	})

	Context("loading from make list", func() {
		It("loads individual referenced entity files from make list", func() {
			source := "../fixtures/makelists/basic.yaml"

			errors := make([]file.Error, 0, 0)
			warnings := make([]file.Error, 0, 0)

			entities, errors, warnings := file.ImportEntities(source, errors, warnings)

			Expect(len(entities.Keys)).To(Equal(1))
			Expect(entities.Keys[0].Link).To(Equal("normal_key"))

			Expect(len(entities.Mobs)).To(Equal(1))
			Expect(entities.Mobs[0].Link).To(Equal("mob_link"))

			Expect(len(entities.Doors)).To(Equal(1))
			Expect(entities.Doors[0].Link).To(Equal("door_link"))
		})
	})
})
