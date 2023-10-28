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
}
type ToDo struct {
	NotificationsEnabled bool      `json:"notificationEnabled"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	AlertTime            time.Time `json:"alertTime"`
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
