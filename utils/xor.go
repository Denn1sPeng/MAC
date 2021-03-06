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

import "errors"

func XOR(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("length not equal")
	}
	t := []byte{}

	for i := 0; i < len(a); i++ {
		t = append(t, a[i]^b[i])
	}
	return t, nil
}
