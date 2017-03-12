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

// An InputSource is a fake implementation of input.Source
type InputSource struct {
	imposter.Fake
	pollEventReturnVal tcell.Event
}

// NewInputSource provides a new fake source
func NewInputSource() *InputSource {
	s := InputSource{}
	return &s
}

// PollEvent is a fake method call
func (s *InputSource) PollEvent() tcell.Event {
	s.SetCall("PollEvent")
	return s.pollEventReturnVal
}

// PollEventReturns sets the return value for PollEvent
func (s *InputSource) PollEventReturns(event tcell.Event) {
	s.pollEventReturnVal = event
}
