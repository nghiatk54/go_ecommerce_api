package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// generate hash key
func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// generate random salt
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// hash password
func HashPassword(password, salt string) string {
	saltedPassword := password + salt
	hashPass := sha256.Sum256([]byte(saltedPassword))
	return hex.EncodeToString(hashPass[:])
}
