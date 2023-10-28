package jwt

import "github.com/sirupsen/logrus"

func TestJWT() bool {
	var email string = "example@email.com"

	firstToken, createErr := CreateJWT(email)

	if createErr != nil {
		logrus.Fatal(createErr)
	}
	isValid := Validate(firstToken)

	jwtEMAIL, err := GetEmail(firstToken)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Email: %v", jwtEMAIL)

	return isValid
}
