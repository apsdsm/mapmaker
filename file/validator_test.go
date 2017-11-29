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
	"github.com/apsdsm/mapmaker/file"

	"github.com/apsdsm/mapmaker/formats/placeholder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapValidator", func() {

	Context("entity rules", func() {
		It("returns error if map references mob that doesn't exist", func() {
			source := "../fixtures/maps/map_with_one_mob.map"

			meta, dungeon := file.ImportMap(source)
			entities := placeholder.NewEntityCollection()

			errors, _ := file.ValidatePlaceholders(meta, dungeon, &entities)

			Expect(len(errors)).To(Equal(1))
			Expect(errors[0].LineNumber).To(Equal(1))
			Expect(errors[0].Message).To(Equal("mob_link is not defined by any entity."))
		})

		Context("door rules", func() {
			It("returns error if door is referenced that doesn't exist", func() {
				source := "../fixtures/maps/map_with_one_door.map"

				meta, dungeon := file.ImportMap(source)
				entities := placeholder.NewEntityCollection()

				errors, _ := file.ValidatePlaceholders(meta, dungeon, &entities)

				Expect(len(errors)).To(Equal(1))
				Expect(errors[0].LineNumber).To(Equal(1))
				Expect(errors[0].Message).To(Equal("door_link is not defined by any entity."))
			})

			It("returns warning if door uses a key that doesn't exist", func() {
				source := "../fixtures/maps/map_with_one_door.map"
				entFile := "../fixtures/entities/door.door.yaml"

				meta, dungeon := file.ImportMap(source)
				entities := placeholder.NewEntityCollection()
				errors := make([]file.Error, 0, 0)
				warnings := make([]file.Error, 0, 0)
				errors, warnings = file.AddEntityToCollection(entFile, &entities, errors, warnings)

				errors, warnings = file.ValidatePlaceholders(meta, dungeon, &entities)

				Expect(len(warnings)).To(Equal(1))
				Expect(warnings[0].LineNumber).To(Equal(1))
				Expect(warnings[0].Message).To(Equal("door_link has no matching key entity."))
			})
		})
	})

	Context("Map rules", func() {
		It("returns error if there is no start position", func() {
			mapFile := "../fixtures/maps/empty_map.map"

			meta, dungeon := file.ImportMap(mapFile)

			entities := placeholder.NewEntityCollection()

			errors, _ := file.ValidatePlaceholders(meta, dungeon, &entities)

			Expect(len(errors)).To(Equal(1))
			Expect(errors[0].LineNumber).To(Equal(-1))
			Expect(errors[0].Message).To(Equal("map has no start position."))
		})
	})

	Context("Meta rules", func() {
		It("returns warning if there is no name", func() {
			mapFile := "../fixtures/maps/identity.map"
			meta, dungeon := file.ImportMap(mapFile)
			entities := placeholder.NewEntityCollection()

			_, warnings := file.ValidatePlaceholders(meta, dungeon, &entities)

			Expect(len(warnings)).To(Equal(1))
			Expect(warnings[0].LineNumber).To(Equal(-1))
			Expect(warnings[0].Message).To(Equal("map has no name."))
		})
	})
})
