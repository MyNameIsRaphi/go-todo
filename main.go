package main

import (
	"net/http"
	"os"
	"time"
	"todo/database"
	"todo/jwt"
	"todo/routes"

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
	
	setupLogger() // init logger

	works := jwt.TestJWT() // test if JWT works
	if works {
		logrus.Info("JWT Works")
	} else {
		logrus.Fatal("JWT not working")
	}
	err := database.ConnectDB() // connect database
	checkError(err)
	http.HandleFunc("/", routes.Middleware)
	checkError(http.ListenAndServe("localhost:8080", nil))

}

func checkError(err error) {
	if err != nil {
		logrus.WithError(err).Warn("Error occurred during operation")
	}
}
