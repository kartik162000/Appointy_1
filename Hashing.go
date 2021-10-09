package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func genRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func hashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)

	var sha512Hasher = sha512.New()

	salt = []byte{4, 17, 178, 104, 25, 207, 47, 37, 106, 40, 231, 24, 174, 36, 12, 104}
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}
