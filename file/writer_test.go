package file_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/maps"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapWriter", func() {
	It("outputs map to json", func() {
		outputFilePath := "../fixtures/__test_output.json"

		_ = os.Remove(outputFilePath)

		m := maps.NewMap(5, 5)

		file.Write(m, outputFilePath)

		infile, _ := ioutil.ReadFile(outputFilePath)
		var m2 maps.Map
		json.Unmarshal(infile, &m2)

		Expect(reflect.DeepEqual(*m, m2)).To(BeTrue())
	})
})
