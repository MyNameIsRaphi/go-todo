package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(word string) (string, error) {
	hashedBytes, hashError := bcrypt.GenerateFromPassword([]byte(word), 10)
	return string(hashedBytes), hashError
}

func Compare(word, hashedWord string) bool {
	compareError := bcrypt.CompareHashAndPassword([]byte(hashedWord), []byte(word))
	return compareError == nil
}
