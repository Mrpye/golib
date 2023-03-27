// Package for working with encryption
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// createHash It takes a string, converts it to a byte array, and then creates a hash of that byte array
// and returns the hash as a string
// - key: the string to hash
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// encrypt It takes a byte array and a passphrase, creates a hash from the passphrase, creates a cipher from
// the hash, creates a nonce, and then encrypts the data with the nonce
// - data: the byte array to encrypt
// - passphrase: the passphrase to use to encrypt the data
// - returns: the encrypted byte array
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

// decrypt It takes a byte array and a passphrase, creates a hash from the passphrase, uses the hash as a key
// to decrypt the byte array, and returns the decrypted byte array
// - data: the byte array to decrypt
// - passphrase: the passphrase to use to decrypt the data
// - returns: the decrypted byte array
// - returns: an error if there is one
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

// Base64EncString It takes a string or []byte, converts it to a byte array, then encodes it to a base64 string
// and returns the base64 string
// - value: the string or []byte to encode
// - returns: the base64 encoded string
func Base64EncString(value interface{}) (string, error) {
	switch v := value.(type) {
	case []byte:
		return b64.StdEncoding.EncodeToString(v), nil
	case string:
		return b64.StdEncoding.EncodeToString([]byte(v)), nil
	default:
		return "", fmt.Errorf("value must be a string or []byte")
	}
}

// Base64DecString 	It takes a string or []byte, converts it to a byte array, then decodes it from a base64 string
// - value: the string to decode
// - returns: the decoded string
// - returns: an error if there is one
func Base64DecString(value interface{}) (string, error) {
	var sDec []byte
	var err error
	switch v := value.(type) {
	case []byte:
		sDec, err = b64.StdEncoding.DecodeString(string(v))
		if err != nil {
			return string(v), err
		}
	case string:
		sDec, err = b64.StdEncoding.DecodeString(string(v))
		if err != nil {
			return string(v), err
		}
	default:
		return "", fmt.Errorf("value must be a string or []byte")
	}
	return string(sDec), nil
}

// EncryptString takes a string or []byte, encrypts it, encodes it to base64, and returns the encrypted string
// - value: the string or []byte to encrypt
// - pass_word: the passphrase to use to encrypt the data
// - returns: the encrypted string
// - returns: an error if there is one
func EncryptString(value interface{}, pass_word string) (string, error) {
	var sDec []byte
	var err error
	switch v := value.(type) {
	case []byte:
		sDec, err = encrypt(v, pass_word)
		if err != nil {
			return string(v), err
		}
	case string:
		sDec, err = encrypt([]byte(v), pass_word)
		if err != nil {
			return string(v), err
		}
	default:
		return "", fmt.Errorf("value must be a string or []byte")
	}
	sEnc := b64.StdEncoding.EncodeToString(sDec)
	return sEnc, nil
}

// DecryptString takes a string or []byte, decodes it from base64, decrypts it, and returns the decrypted string
// - value: the string or []byte to decrypt
// - pass_word: the passphrase to use to decrypt the data
// - returns: the decrypted string
// - returns: an error if there is one
func DecryptString(value interface{}, pass_word string) (string, error) {
	var sDec []byte
	var err error
	switch v := value.(type) {
	case []byte:
		sDec, err = b64.StdEncoding.DecodeString(string(v))
		if err != nil {
			return string(v), err
		}
	case string:
		sDec, err = b64.StdEncoding.DecodeString(v)
		if err != nil {
			return string(v), err
		}
	default:
		return "", fmt.Errorf("value must be a string or []byte")
	}

	data, err := decrypt(sDec, pass_word)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
