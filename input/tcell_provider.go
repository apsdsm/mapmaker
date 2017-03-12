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

package input

import (
	"github.com/gdamore/tcell"
)

// Provider converts tcell input to application input
type TCellProvider struct {
	source Source
}

// NewProvider cretes a new Provider instance
func NewTcellProvider(source Source) *TCellProvider {
	c := TCellProvider{
		source: source,
	}
	return &c
}

// GetInput will return any input collected this update
func (c *TCellProvider) GetInput() int {
	event := c.source.PollEvent()

	switch e := event.(type) {

	case *tcell.EventKey:
		return c.ToCode(e)
	}

	return -1
}

// ToCode converts a tcell key event to an input code
func (c *TCellProvider) ToCode(event *tcell.EventKey) int {

	// if user pressed q
	if event.Key() == tcell.KeyRune && string(event.Rune()) == "q" {
		return Quit
	}

	return -1
}
