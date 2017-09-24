package file_test

import (
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/placeholders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapReader", func() {
	It("inputs map from json", func() {
		inputFilePath := "../fixtures/simple_map.json"

		m := file.Read(inputFilePath)

		file, _ := ioutil.ReadFile(inputFilePath)
		var m2 placeholders.Map
		json.Unmarshal(file, &m2)

		Expect(reflect.DeepEqual(*m, m2)).To(BeTrue())
	})
})
