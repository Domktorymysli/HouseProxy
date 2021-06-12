package cypher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(key []byte, iv []byte, msg []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	content := PKCS5Padding(msg, block.BlockSize())
	encrypted := make([]byte, len(content))

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(encrypted, content)

	return encrypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
