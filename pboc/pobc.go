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
	"errors"
	"bytes"
	"Denn1sPeng/MAC/utils"
)

type Pboc struct {
	IV []byte //Initial Vector
}

func (p Pboc) InitIV(iv []byte) error {
	if len(iv) != 8 {
		return errors.New("length error")
	}
	p.IV = iv
	return nil
}

func (p Pboc) GenDesMac(key, data []byte) ([]byte, error) {
	block_num := len(data)/8 + 1
	block_byte_len := block_num * 8
	data = append(data, 0x80)
	block_tail_len := block_byte_len - len(data)
	data = append(data, bytes.Repeat([]byte{0}, block_tail_len)...)

	sign_data := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	params := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	if len(p.IV) == 0 {
		copy(sign_data, p.IV)
	}

	sign_data, err := utils.XOR(sign_data, data[:8])
	if err != nil {
		return nil, err
	}

	params, err = utils.DesEncrypt(sign_data, key)
	if err != nil {
		return nil, err
	}

	//encrypt the rest
	for i := 1; i < block_num; i++ {
		tempt, err := utils.XOR(params, data[i*8:(i+1)*8])
		if err != nil {
			return nil, err
		}

		params, err = utils.DesEncrypt(tempt, key)
		if err != nil {
			return nil, err
		}
	}

	return params[:len(key)], nil

}

func (p Pboc) Gen3DesMac(key, data []byte) ([]byte, error) {
	block_num := len(data)/8 + 1
	block_byte_len := block_num * 8
	data = append(data, 0x80)
	block_tail_len := block_byte_len - len(data)
	data = append(data, bytes.Repeat([]byte{0}, block_tail_len)...)

	sign_data := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	params := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	if len(p.IV) == 0 {
		copy(sign_data, p.IV)
	}

	sign_data, err := utils.XOR(sign_data, data[:8])
	if err != nil {
		return nil, err
	}

	front_key := key[:8]
	for i := 1; i < block_num; i++ {
		params, err = utils.DesEncrypt(sign_data, front_key)
		if err != nil {
			return nil, err
		}
		sign_data, err = utils.XOR(params, data[i*8:(i+1)*8])
	}

	sign_data, err = utils.Des3Encrypt(sign_data, key)
	if err != nil {
		return nil, err
	}

	return sign_data, nil

}
