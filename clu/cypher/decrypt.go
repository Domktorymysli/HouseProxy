package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func Decrypt(key []byte, iv []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	decrypted := make([]byte, len(ciphertext))

	if len(ciphertext)%aes.BlockSize != 0 {
		return make([]byte, 0), errors.New("crypto/cipher: input not full blocks")
	}

	mode.CryptBlocks(decrypted, ciphertext)

	return PKCS5Trimming(decrypted), nil
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
