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
	"github.com/apsdsm/mapmaker/placeholders"
)

var _ = Describe("EntityImporter", func() {
	It("loads a yaml file of mob", func() {
		entities := placeholders.NewEntityCollection()
		file.AddEntityToCollection("../fixtures/entities/mob.mob.yaml", entities)

		Expect(len(entities.Mobs)).To(Equal(1))
		Expect(entities.Mobs[0].Name).To(Equal("Mob Name"))
		Expect(entities.Mobs[0].Link).To(Equal("mob_link"))
		Expect(entities.Mobs[0].Prot).To(Equal("mob_prot"))

		file.AddEntityToCollection("../fixtures/entities/mob_prot.mob.yaml", entities)

		Expect(len(entities.Mobs)).To(Equal(2))
		// TODO: insert prototype expectations
	})

	It("loads a yaml file of door", func() {
		source := "../fixtures/entities/door.door.yaml"

		entities := placeholders.NewEntityCollection()
		file.AddEntityToCollection(source, entities)

		Expect(len(entities.Doors)).To(Equal(1))
		Expect(entities.Doors[0].Link).To(Equal("door_link"))
		Expect(entities.Doors[0].Locked).To(Equal(true))
		Expect(entities.Doors[0].Key).To(Equal("door_key"))
		Expect(entities.Doors[0].OnTry).To(Equal("door_ontry"))
	})

	It("loads a yaml file of key", func() {
		source := "../fixtures/entities/key.key.yaml"

		entities := placeholders.NewEntityCollection()
		file.AddEntityToCollection(source, entities)

		Expect(len(entities.Keys)).To(Equal(1))
		Expect(entities.Keys[0].Name).To(Equal("Normal Key"))
		Expect(entities.Keys[0].Link).To(Equal("normal_key"))
		Expect(entities.Keys[0].Desc).To(Equal("A normal key"))
	})

	It("loads individual referenced entity files from make list", func() {
		source := "../fixtures/makelists/basic.yaml"

		entities := file.ImportEntities(source)

		Expect(len(entities.Keys)).To(Equal(1))
		Expect(entities.Keys[0].Link).To(Equal("normal_key"))

		Expect(len(entities.Mobs)).To(Equal(1))
		Expect(entities.Mobs[0].Link).To(Equal("mob_link"))

		Expect(len(entities.Doors)).To(Equal(1))
		Expect(entities.Doors[0].Link).To(Equal("door_link"))
	})
})
