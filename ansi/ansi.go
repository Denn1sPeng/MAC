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
	"Denn1sPeng/MAC/utils"
	"errors"
)

type Ansi struct {
	IV []byte //Initial Vector
}

func (m Ansi) InitIV(iv []byte) error {
	if len(iv) != 8 {
		return errors.New("length error")
	}
	m.IV = iv
	return nil
}

func (m Ansi) GenX99Mac(key, data []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("key lenght error")
	}
	var err error = nil

	data_length := len(data)
	//fullfil the last block with 0
	if data_length%8 != 0 {
		for i := 0; i < 8-data_length%8; i++ {
			data = append(data, 0)
		}
	}

	sign_data := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	params := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	if len(m.IV) == 0 {
		copy(sign_data, m.IV)
	}

	for i := 0; i < len(data)/8; i++ {
		params, err = utils.XOR(data[i*8:(i+1)*8], sign_data)
		if err != nil {
			return nil, err
		}

		sign_data, err = utils.DesEncrypt(params, key)
		if err != nil {
			return nil, err
		}

	}

	return sign_data, nil
}

func (m Ansi) GenX919Mac(key, data []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("key length error.")
	}

	mac_key_left_bytes := key[0:8]
	mac_key_right_bytes := key[8:16]

	sign99, err := m.GenX99Mac(mac_key_left_bytes, data)
	if err != nil {
		return nil, err
	}

	params, err := utils.DesDecrypt(sign99, mac_key_right_bytes)
	if err != nil {
		return nil, err
	}

	sign_data, err := utils.DesEncrypt(params, mac_key_left_bytes)
	if err != nil {
		return nil, err
	}

	return sign_data, nil
}
