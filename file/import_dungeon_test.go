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

func checkLevelHasEnoughCells(level *placeholder.Map) {
	Expect(len(level.Grid)).To(Equal(level.Width))

	for i := range level.Grid {
		Expect(len(level.Grid[i])).To(Equal(level.Height))
	}
}

var _ = Describe("MapImporter", func() {
	It("loads the metadata at the top of a map file into a placeholder", func() {
		source := "../fixtures/maps/meta_only.map"

		meta, _ := file.ImportDungeon(source)

		Expect(meta.Link).To(Equal("prison"))
		Expect(meta.Name).To(Equal("The Jovian Prison"))
		Expect(meta.Desc).To(Equal("A gloomy building hidden deep in the Jovian woods."))
	})

	It("loads the dungeon after metadata into a placeholder", func() {
		source := "../fixtures/maps/meta_and_map.map"

		_, level := file.ImportDungeon(source)

		Expect(level.Width).To(Equal(3))
		Expect(level.Height).To(Equal(3))

		checkLevelHasEnoughCells(level)
	})

	It("loads the metadata after each line of map content", func() {
		source := "../fixtures/maps/meta_and_annotated_map.map"

		_, level := file.ImportDungeon(source)

		Expect(level.Width).To(Equal(8), "width should be 8")
		Expect(level.Height).To(Equal(5), "height should be 4")

		checkLevelHasEnoughCells(level)

		Expect(level.Grid[2][1].Type).To(Equal("mob"))
		Expect(level.Grid[2][1].Link).To(Equal("mob_link"))

		Expect(level.Grid[2][2].Type).To(Equal("door"))
		Expect(level.Grid[2][2].Link).To(Equal("door_link"))

		Expect(level.Grid[2][3].Type).To(Equal("waypoint"))
		Expect(level.Grid[2][3].Link).To(Equal("waypoint1"))
	})

	It("assigns a start position if it finds one", func() {
		source := "../fixtures/maps/meta_and_annotated_map.map"

		_, dungeon := file.ImportDungeon(source)

		Expect(dungeon.StartPosition).ToNot(BeNil())

		Expect(dungeon.StartPosition.X).To(Equal(4))
		Expect(dungeon.StartPosition.Y).To(Equal(3))
	})

	It("sets the rune value to ' ' if cell contains entity", func() {
		source := "../fixtures/maps/meta_and_annotated_map.map"

		_, dungeon := file.ImportDungeon(source)

		// start position
		Expect(dungeon.Grid[4][3].Rune).To(Equal(' '))
	})
})