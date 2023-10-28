package main

import (
	"os"
	"time"
	"todo/database"
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

	connectionError := database.ConnectDB() // connect database

	if connectionError != nil {
		logrus.Fatal(connectionError)
	} else {
		logrus.Info("Successfully connected to db")
	}

	user := types.User{
		Email:                "raphaelkaeser05@gmail.com",
		Password:             "12344242343",
		NotificationsGranted: false,
	}
	checkError(database.AddUser(user))

	foundUser, err := database.GetUser(user.Email)

	checkError(err)

	logrus.Infof("Found user: %v", foundUser)

}

func checkError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
