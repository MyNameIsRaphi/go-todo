package database

import (
	"context"
	"todo/encrypt"
	"todo/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Uri string = "mongodb://localhost:27017"
var ServerAPI = options.ServerAPI(options.ServerAPIVersion1)
var Option = options.Client().ApplyURI(Uri).SetServerAPIOptions(ServerAPI)
var pOptions = &Option

const Database string = "todo"
const Collection string = "todo"

func ConnectDB() (*mongo.Client, error) {
	var Client, connectionError = mongo.Connect(context.TODO(), *pOptions)

	return Client, connectionError

}

func AddUser(client *mongo.Client, user types.User) error {
	hashedPassword, hashError := encrypt.Hash(user.Password)
	if hashError != nil {
		panic(hashError)
	}

	user.Password = hashedPassword
	_, insertError := client.Database(Database).Collection(Collection).InsertOne(context.TODO(), user)
	return insertError
}
func DisconnectDB(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}
