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
)

// An InputProvider is a fake implementation of input.Provider
type InputProvider struct {
	imposter.Fake
	getInputReturnValue int
}

// NewInputProvider returns a new InputProvider
func NewInputProvider() *InputProvider {
	i := InputProvider{}
	return &i
}

// GetInput fakes a call to this method
func (i *InputProvider) GetInput() int {
	i.SetCall("GetInput")
	return i.getInputReturnValue
}

// GetInputReturns sets the return value for GetInput
func (i *InputProvider) GetInputReturns(returns int) {
	i.getInputReturnValue = returns
}
