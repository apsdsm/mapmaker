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

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapImporter", func() {
	It("loads the metadata at the top of a map file into a placeholder", func() {
		meta, _ := file.ImportDungeon("../fixtures/maps/meta_only.map")

		Expect(meta.Link).To(Equal("prison"))
		Expect(meta.Name).To(Equal("The Jovian Prison"))
		Expect(meta.Desc).To(Equal("A gloomy building hidden deep in the Jovian woods."))
	})

	It("loads the dungeon after metadata into a placeholder", func() {
		_, level := file.ImportDungeon("../fixtures/maps/meta_and_map.map")

		Expect(level.Width).To(Equal(3))
		Expect(level.Height).To(Equal(3))
		Expect(len(level.Grid)).To(Equal(level.Width))

		for i := range level.Grid {
			Expect(len(level.Grid[i])).To(Equal(level.Height))
		}
	})

	It("loads the metadata after each line of map content", func() {
		_, level := file.ImportDungeon("../fixtures/maps/meta_and_annotated_map.map")

		Expect(level.Grid[0][1].Type).To(Equal("mob"))
		Expect(level.Grid[0][1].Link).To(Equal("mob1"))

		Expect(level.Grid[0][2].Type).To(Equal("door"))
		Expect(level.Grid[0][2].Link).To(Equal("door1"))

		Expect(level.Grid[0][3].Type).To(Equal("waypoint"))
		Expect(level.Grid[0][3].Link).To(Equal("waypoint1"))

		Expect(level.Grid[0][4].Type).To(Equal("item"))
		Expect(level.Grid[0][4].Link).To(Equal("item1"))
	})

	It("assigns a start position if it finds one", func() {
		_, dungeon := file.ImportDungeon("../fixtures/maps/map_with_start.map")

		Expect(dungeon.StartPosition).ToNot(BeNil())
		Expect(dungeon.StartPosition.X).To(Equal(1))
		Expect(dungeon.StartPosition.Y).To(Equal(1))
	})

	It("sets the rune value to ' ' if cell contains entity", func() {
		_, dungeon := file.ImportDungeon("../fixtures/maps/map_with_one_mob.map")

		Expect(dungeon.Grid[0][0].Rune).To(Equal(' '))
	})
})
