package encrypt

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"github.com/sirupsen/logrus"
)

func Hash(word string) (string, error) {
	var h hash.Hash = sha512.New()

	_, err := h.Write([]byte(word))
	var hashedBytes []byte = h.Sum(nil)

	var hashedWord string = hex.EncodeToString(hashedBytes)
	logrus.Infof("Hash:%v ", hashedWord)
	return hashedWord, err
}

func Compare(word, hashedWord string) bool {
	hWord, err := Hash(word)
	if err != nil {
		logrus.WithError(err).Warn("Couldn't hash word")
		return false
	}
	return hWord == hashedWord
}
