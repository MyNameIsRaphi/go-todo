package jwt

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"todo/config"
)

var key string = config.Config.JWT_KEY
var blacklist []string

func CreateJWT(email string) (string, error) {
	header, headerErr := createHeader()

	if headerErr != nil {
		return "", headerErr
	}

	payload, payloadErr := createPayload(email)

	if payloadErr != nil {
		return "", payloadErr
	}

	signature, sigErr := createSignature(header, payload)

	if sigErr != nil {
		return "", sigErr
	}

	return header + "." + payload + "." + signature, nil

}
func AddToBlackList(token string) {
	blacklist = append(blacklist, token)
}
func RemoveFromBlacklist(token string) error {
	var foundIndex int = -1

	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] == token {
			foundIndex = i
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("failed to find token in blacklist")
	}

	blacklist = append(blacklist[:foundIndex], blacklist[foundIndex+1:]...)
	return nil
}
func GetEmail(token string) (string, error) {
	var payload string = strings.SplitAfter(token, ".")[1]
	decodedBytes, decodeErr := decode64(payload)
	if decodeErr != nil {
		return "", decodeErr
	}
	var pl jwtPayload

	unmarshalErr := json.Unmarshal(decodedBytes, &pl)

	if unmarshalErr != nil {
		return "", unmarshalErr
	}
	return pl.Sub, nil
}

func Validate(token string) bool {
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] == token {
			return false
		}
	}

	splitedToken := strings.SplitAfter(token, ".")

	signature := splitedToken[2]
	signature = strings.ReplaceAll(signature, ".", "")
	verifyHeader := splitedToken[0]
	verifyHeader = strings.ReplaceAll(verifyHeader, ".", "")
	verifyPayload := splitedToken[1]
	verifyPayload = strings.ReplaceAll(verifyPayload, ".", "")

	return verify(verifyHeader+"."+verifyPayload, signature)

}

func verify(sig, actSig string) bool {
	var HMAC string = createHMAC(sig)
	return HMAC == actSig
}

func createHeader() (string, error) {
	header := jwtHeader{
		Alg: "HS512",
		Typ: "JWT",
	}
	headerBytes, marshalErr := json.Marshal(header)

	if marshalErr != nil {
		return "", marshalErr
	}
	return encode64(headerBytes), nil

}

func createPayload(email string) (string, error) {
	var exp int64 = time.Now().Add(24 * time.Hour).Unix()
	payload := jwtPayload{
		Exp: exp,
		Sub: email,
	}

	payloadBytes, marshalErr := json.Marshal(payload)

	if marshalErr != nil {
		return "", marshalErr
	}
	return encode64(payloadBytes), nil
}
func createSignature(header, payload string) (string, error) {
	decryptSig := header + "." + payload
	if key == "" {
		return "", fmt.Errorf("No Key declared")

	}
	return createHMAC(decryptSig), nil

}

func createHMAC(s string) string {
	mac := hmac.New(sha512.New, []byte(key))

	mac.Write([]byte(s))
	return hex.EncodeToString(mac.Sum(nil))
}

func encode64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func decode64(s string) ([]byte, error) {
	s = strings.ReplaceAll(s, ".", "")
	decodedBytes, decodeErr := base64.StdEncoding.DecodeString(s)

	return decodedBytes, decodeErr
}

type jwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type jwtPayload struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}
