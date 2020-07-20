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
	"errors"
)

func Des3Encrypt(src, key []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("error key")
	}
	//1:取密钥前8个字节数据采用DES加密
	front8Key := key[:8]
	firstRes, err := DesEncrypt(src, front8Key)
	if err != nil {
		return nil, err
	}
	//2:取密钥后8个字节采用DES解密
	back8Key := key[8:]
	secondRes, err := DesDecrypt(firstRes, back8Key)
	if err != nil {
		return nil, err
	}
	//3:再次用前8个字节的密钥采用DES加密
	last, err := DesEncrypt(secondRes, front8Key)
	if err != nil {
		return nil, err
	}
	return last, nil
}

func Des3Decrypt(src, key []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("error key")
	}
	//1:取密钥前8个字节数据采用DES 解密
	front8Key := key[:8]
	firstRes, err := DesDecrypt(src, front8Key)
	if err != nil {
		return nil, err
	}
	//2:取密钥后8个字节采用DES加密
	back8Key := key[8:]
	secondRes, err := DesEncrypt(firstRes, back8Key)
	if err != nil {
		return nil, err
	}
	//3:再次用前8个字节的密钥采用DES解密
	last, err := DesDecrypt(secondRes, front8Key)
	if err != nil {
		return nil, err
	}
	return last, nil
}
