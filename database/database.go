package database

import (
	"context"
	"fmt"
	"os"
	"todo/encrypt"
	"todo/types"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// set Database and collection name
const Database string = "todo"
const Collection string = "todo"

// load username and password from env for DB
var username = os.Getenv("USERNAME")
var password = os.Getenv("PASSWORD")

var Uri string = fmt.Sprintf("mongodb://%v:%v@localhost:27017", username, password)
var ServerAPI = options.ServerAPI(options.ServerAPIVersion1)
var Option = options.Client().ApplyURI(Uri).SetServerAPIOptions(ServerAPI)
var client *mongo.Client
var todoCollection *mongo.Collection

func ConnectDB() {

	cl, connectionError := mongo.Connect(context.TODO(), Option)
	if connectionError != nil {
		logrus.WithError(connectionError).Fatal("Couldn't connect to database")
	}
	client = cl
	todoCollection = client.Database(Database).Collection(Collection)
}

func AddUser(user types.User) error {
	if Exists(user.Email) {
		return fmt.Errorf("User with this email address already exists")
	}
	hashedPassword, hashError := encrypt.Hash(user.Password)
	if hashError != nil {
		return hashError
	}

	user.Password = hashedPassword
	_, insertError := todoCollection.InsertOne(context.TODO(), user)
	return insertError
}
func DisconnectDB(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}
func Exists(email string) bool {
	filter := bson.M{
		"email": email,
	}
	result := todoCollection.FindOne(context.TODO(), filter)

	return result.Err() == nil
}
func GetUser(email string) (types.User, error) {

	filter := bson.M{"email": email}
	var result mongo.SingleResult = *todoCollection.FindOne(context.TODO(), filter)

	var foundUser types.User

	decodeError := result.Decode(&foundUser)

	return foundUser, decodeError

}

func CheckCreditentials(email, password string) bool {
	user, foundError := GetUser(email)

	if foundError != nil {
		logrus.Warn("Couldn't find user")
		return false
	}

	return encrypt.Compare(password, user.Password)
}
