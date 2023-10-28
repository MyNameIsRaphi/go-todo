package main

import (
	"os"
	"time"
	"todo/database"
	"todo/jwt"
	"todo/types"

	"github.com/sirupsen/logrus"
)

func setupLogger() {
	logrus.SetOutput(os.Stderr)

	var textFormatter logrus.TextFormatter = logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: time.RFC3339Nano,
		FullTimestamp:   true,
		PadLevelText:    true,
		DisableQuote:    false,
	}
	logrus.SetFormatter(&textFormatter)
}

func main() {

	setupLogger() // inistalize logger

	works := jwt.TestJWT() // test if JWT works
	if works {
		logrus.Info("JWT Works")
	} else {
		logrus.Fatal("JWT not working")
	}
	connectionError := database.ConnectDB() // connect database

	if connectionError != nil {
		logrus.Fatal(connectionError)
	} else {
		logrus.Info("Successfully connected to db")
	}

	user := types.User{
		Email:                "raphaelkdfdfaeser05@gmail.com",
		Password:             "12344242343",
		NotificationsGranted: false,
	}
	checkError(database.AddUser(user))

	foundUser, err := database.GetUser(user.Email)

	checkError(err)

	logrus.Infof("Found user: %v", foundUser)

	var isValid bool = database.CheckCreditentials(user.Email, user.Password)

	logrus.Info(isValid)

}

func checkError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
