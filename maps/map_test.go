//    Copyright 2017 Nick del Pozo
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package maps_test

import (
	"github.com/apsdsm/mapmaker/maps"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map", func() {

	It("makes a new map with specified width and height", func() {
		m := maps.NewMap(10, 5)

		Expect(m.Width).To(Equal(10))
		Expect(m.Height).To(Equal(5))

		Expect(len(m.Grid)).To(Equal(10))

		for i := range m.Grid {
			Expect(len(m.Grid[i])).To(Equal(5))
		}
	})
})
