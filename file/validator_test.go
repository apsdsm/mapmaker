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

	var (
		errors   []file.Error
		warnings []file.Error
	)

	BeforeEach(func() {

		// empty errors/warnings
		errors = make([]file.Error, 0, 0)
		warnings = make([]file.Error, 0, 0)
	})

	It("returns error if cell contains mob that doesn't exist", func() {
		// meta
		meta := placeholder.Meta{}
		meta.Name = "foo dungeon"

		// dungeon
		dungeon := placeholder.NewMap(1, 1)
		dungeon.Grid[0][0] = placeholder.EmptyCell()
		dungeon.Grid[0][0].Annotated = true
		dungeon.Grid[0][0].Link = "mob"
		dungeon.Grid[0][0].Type = "mob"
		dungeon.StartPosition = &placeholder.Position{0, 0}

		// empty entity collection
		entities := placeholder.NewEntityCollection()

		errors, _ = file.ValidatePlaceholders(&meta, dungeon, &entities)

		Expect(len(errors)).To(Equal(1))
		Expect(errors[0].LineNumber).To(Equal(1))
		Expect(errors[0].Message).To(Equal("mob is not defined by any entity."))
	})

	It("returns error if cell contains door that doesn't exist", func() {
		// meta
		meta := placeholder.Meta{}
		meta.Name = "foo dungeon"

		// dungeon
		dungeon := placeholder.NewMap(1, 1)
		dungeon.Grid[0][0] = placeholder.EmptyCell()
		dungeon.Grid[0][0].Annotated = true
		dungeon.Grid[0][0].Link = "door_link"
		dungeon.Grid[0][0].Type = "door"
		dungeon.StartPosition = &placeholder.Position{0, 0}

		// empty entity collection
		entities := placeholder.NewEntityCollection()

		errors, _ = file.ValidatePlaceholders(&meta, dungeon, &entities)

		Expect(len(errors)).To(Equal(1))
		Expect(errors[0].LineNumber).To(Equal(1))
		Expect(errors[0].Message).To(Equal("door_link is not defined by any entity."))
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
