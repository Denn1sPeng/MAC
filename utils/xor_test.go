//Copyright 2020 Denn1sPeng
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package utils

import (
	"testing"
	"fmt"
)

func TestXOR(t *testing.T) {

	a := []byte{0x8, 0x0, 0x8, 0x0}
	b := []byte{0x7, 0x1, 0x5, 0x3}

	c, err := XOR(a, b)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	fmt.Printf("%X", c)
}
