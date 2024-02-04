package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func AesEncryptText(raw []byte, passphrase string) ([]byte, error) {
	aesBlock, err := aes.NewCipher([]byte(mdHashing(passphrase)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, raw, nil)

	return cipheredText, nil
}

func AesDecryptText(ciphered []byte, keyPhrase string) ([]byte, error) {
	hashedPhrase := mdHashing(keyPhrase)
	aesBlock, err := aes.NewCipher([]byte(hashedPhrase))
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]
	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	return originalText, nil
}
