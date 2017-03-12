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

package fakes

import (
	"github.com/apsdsm/imposter"
	"github.com/gdamore/tcell"
)

// A Screen is a fake implementatin of the Screen
type Screen struct {
	imposter.Fake
}

// NewScreen provides a new fake screen
func NewScreen() *Screen {
	s := Screen{}
	return &s
}

// signature for SetContent
type setContentSig struct {
	X, Y  int
	Mainc rune
	Combc []rune
	Style tcell.Style
}

// SetContent emulates call to this method
func (s *Screen) SetContent(x int, y int, mainc rune, combc []rune, style tcell.Style) {
	s.SetCall("SetContent", setContentSig{x, y, mainc, combc, style})
}

// Size emulates call to this method
func (s *Screen) Size() (int, int) {
	s.SetCall("Size")
	return 0, 0
}
