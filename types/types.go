package types

import (
	"time"
)

type User struct {
	Email                string `json:"email" validate:"required email"`
	Password             string `json:"password" validate:"required min:7"`
	NotificationsGranted bool   `json:"notificationsGranted"`
	//	ID	primitive.ObjectID `bson:"_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ID        string `json:"id"`
	ToDos     []ToDo `json:"toDos"`
}
type ToDo struct {
	NotificationsEnabled bool      `json:"notificationEnabled"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	AlertTime            time.Time `json:"alertTime"`
	Date                 time.Time `json:"date"`
	Status               bool      `json:"status"`
}

type ReqgisterRequest struct {
	Email     string `json:"email"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
