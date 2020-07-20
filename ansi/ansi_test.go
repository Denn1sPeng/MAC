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

package ansi

import (
	"testing"
	"encoding/hex"
	"strings"
)

func TestGenX99Mac(t *testing.T) {
	want := "2C8989B7028CE438"
	key, _ := hex.DecodeString("F76213D3FD26CBAD")
	data, _ := hex.DecodeString("0200722044C020C0909118628882001030600201000000000000050000121218173800022012340210000636628882001030600201D120452010000595004757313030303237383838313532303030303030303839313536D3078ED63706EDBD0230310011220000010000")
	ansi := Ansi{}
	result, err := ansi.GenX99Mac(key, data)
	if err != nil {
		t.Fatalf("error:%s", err.Error())
		return
	}

	mac := strings.ToUpper(hex.EncodeToString(result))
	if mac != want {
		t.Fatalf("Mac not match!Mac:%s Want:%s", mac, want)
	}
}

func TestGenX919Mac(t *testing.T) {

	want, _ := hex.DecodeString("C330D183D7D1564F")
	key, _ := hex.DecodeString("F76213D3FD26CBAD0BC2755DB0D5F810")
	data, _ := hex.DecodeString("0200722044C020C0909118628882001030600201000000000000050000121218173800022012340210000636628882001030600201D120452010000595004757313030303237383838313532303030303030303839313536D3078ED63706EDBD0230310011220000010000")
	ansi := Ansi{}
	result, err := ansi.GenX919Mac(key, data)
	if err != nil {
		t.Fatalf("error:%s", err.Error())
		return
	}

	mac := hex.EncodeToString(result)
	t.Logf("result:%s", strings.ToUpper(mac))

	if mac != hex.EncodeToString(want) {
		t.Fatalf("Mac not match!")
	}
}
