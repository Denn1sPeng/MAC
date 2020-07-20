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

package pboc

import (
	"testing"
	"fmt"
)

func TestPboc_GenDesMac(t *testing.T) {
	key := []byte{0x6F, 0x71, 0x5F, 0xB2, 0x8C, 0x52, 0x33, 0x14}
	data := []byte{0x00, 0x00, 0x00, 0x0A, 0x06, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x00, 0x00, 0x00, 0x1D, 0x20, 0x15, 0x04, 0x14, 0x19, 0x43, 0x00}
	want := []byte{0x27, 0xCF, 0x46, 0xD9, 0x3A, 0x1C, 0x31, 0x8B}

	p := Pboc{}

	mac, err := p.GenDesMac(key, data)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	src := fmt.Sprintf("%X", mac)
	dst := fmt.Sprintf("%X", want)

	if src != dst {
		t.Errorf("Mac not match!")
		return
	}

	fmt.Printf("Mac:%s wanted:%s", src, dst)
}

func TestPboc_Gen3DesMac(t *testing.T) {
	key := []byte{0x57, 0xE0, 0x90, 0x59, 0x25, 0x3A, 0xDC, 0x76, 0xDC, 0x01, 0x55, 0x43, 0x78, 0x1E, 0x00, 0x6A}
	data := []byte{0x00, 0x00, 0xBC, 0xAB, 0x00, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00}
	want := []byte{0x4B, 0xA9, 0x58, 0xD1, 0x78, 0x98, 0xB2, 0x1D}
	p := Pboc{}

	mac, err := p.Gen3DesMac(key, data)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	src := fmt.Sprintf("%X", mac)
	dst := fmt.Sprintf("%X", want)

	if src != dst {
		t.Errorf("Mac not match!")
		return
	}

	fmt.Printf("Mac:%s wanted:%s", src, dst)

}
