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
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/formats/placeholder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapReader", func() {
	It("inputs map from json", func() {
		inputFilePath := "../fixtures/simple_map.json"

		m := file.Read(inputFilePath)

		file, _ := ioutil.ReadFile(inputFilePath)
		var m2 placeholder.Map
		json.Unmarshal(file, &m2)

		Expect(reflect.DeepEqual(*m, m2)).To(BeTrue())
	})
})
