// Package for working with encryption
package encrypt

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

// It takes a string, converts it to a byte array, and then creates a hash of that byte array
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// It takes a byte array and a passphrase, creates a hash from the passphrase, creates a cipher from
// the hash, creates a nonce, and then encrypts the data with the nonce
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

// It takes a byte array and a passphrase, creates a hash from the passphrase, uses the hash as a key
// to decrypt the byte array, and returns the decrypted byte array
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

// It takes a string or []byte, converts it to a byte array, then encodes it to a base64 string
func Base64EncString(value interface{}) string {
	switch v := value.(type) {
	case []byte:
		return b64.StdEncoding.EncodeToString(v)
	case string:
		return b64.StdEncoding.EncodeToString([]byte(v))
	}
	return ""
}

// It takes a string, decodes it from base64, and returns the decoded string
func Base64DecString(value string) string {
	sDec, err := b64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Print(err)
		return "Decode Error"
	}
	return string(sDec)
}

// EncryptString takes a string and a password and returns a base64 encoded string
func EncryptString(value string, pass_word string) (string, error) {
	data, err := encrypt([]byte(value), pass_word)
	if err != nil {
		return value, err
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc, nil
}

// DecryptString takes a string, decodes it from base64, decrypts it, and returns the decrypted string
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
