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

package file

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// readBytes will read and return the bytes from a file path and panic if it can't
func ReadBytes(filePath string) *[]byte {
	file, err := os.Open(filePath)

	defer file.Close()

	if err != nil {
		panic("error while opening file: " + err.Error())
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		panic("error while reading data from file: " + err.Error())
	}

	return &bytes
}

// absPath will run abspath on a path string and panic if it can't
func AbsPath(filePath string) string {
	absPath, err := filepath.Abs(filePath)

	if err != nil {
		panic("couldn't make file path")
	}

	return absPath
}

// unmarshalYaml will try to unmarshal yaml from a byte array and pannic if it can't
func UnmarshalYaml(in []byte, out interface{}) {
	err := yaml.Unmarshal(in, out)

	if err != nil {
		panic("error unmarshalling yaml")
	}
}
