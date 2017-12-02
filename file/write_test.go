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
	"os"
	"reflect"

	"github.com/apsdsm/mapmaker/file"

	"github.com/apsdsm/mapmaker/formats/json_format"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapWriter", func() {
	It("outputs map to json", func() {
		outputFilePath := "../fixtures/__test_output.json"

		// preemptive cleanup
		_ = os.Remove(outputFilePath)

		m := json_format.NewDungeon(5, 5)

		file.Write(m, outputFilePath)

		infile, _ := ioutil.ReadFile(outputFilePath)
		var m2 json_format.Dungeon
		json.Unmarshal(infile, &m2)

		Expect(reflect.DeepEqual(*m, m2)).To(BeTrue())

		// clean up
		_ = os.Remove(outputFilePath)
	})
})
