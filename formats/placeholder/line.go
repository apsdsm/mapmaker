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

package placeholder

// Line is a placeholder for each line of the structure part of a .map file
type Line struct {
	Cells         []*Cell
	RawAnnoations string
}

// NewLine makes and returns a new line
func NewLine(len int) Line {
	line := Line{
		Cells: make([]*Cell, len),
	}
	return line
}

// Len returns the length of the map line
func (l *Line) Len() int {
	return len(l.Cells)
}
