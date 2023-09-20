package main

import (
	"time"
	"todo/database"
	"todo/types"
)

func main() {
	Client, connectionError := database.ConnectDB()

	if connectionError != nil {
		panic(connectionError)
	}

	user := types.User{
		Email:                "raphaelkaeser05@gmail.com",
		Password:             "123434242423434",
		NotificationsGranted: false,
		TimeRegistered:       time.Now(),
	}
	database.AddUser(Client, user)
}
