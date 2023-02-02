package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/hex"
	"io"
	"log"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphered := gcm.Seal(nonce, nonce, data, nil)
	return ciphered, nil
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, err
	}
	nonce, ciphered := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphered, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// Base64EncString encodes a string to base64
func Base64EncString(value string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(value))
	return sEnc
}

/*func DecodeB64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	base64.StdEncoding.Decode(base64Text, []byte(message))
	return strings.ReplaceAll(string(base64Text), "\x00", "")
}*/

// Base64DecString decodes a base64 string
func Base64DecString(value string) string {
	sDec, err := b64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Print(err)
		return "Decode Error"
	}
	return string(sDec)
}

// EncryptString encrypts a string
func EncryptString(value string, pass_word string) (string, error) {
	data, err := encrypt([]byte(value), pass_word)
	if err != nil {
		return value, err
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc, nil
}

// DecryptString decrypts a string
func DecryptString(value string, pass_word string) (string, error) {
	sDec, err := b64.StdEncoding.DecodeString(value)
	if err != nil {
		return value, err
	}

	data, err := decrypt(sDec, pass_word)
	if err != nil {
		return value, err
	}
	return string(data), nil
}
