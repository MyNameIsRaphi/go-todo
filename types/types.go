package types

import (
	"time"
)

type User struct {
	Email                string    `json:"email" validate:"required email"`
	Password             string    `json:"password" validate:"required min:7"`
	NotificationsGranted bool      `json:"notificationsGranted"`
	TimeRegistered       time.Time `json:"timeRegistered"`
}
type ToDo struct {
	NotificationsEnabled bool      `json:"notificationEnabled"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	AlertTime            time.Time `json:"alertTime"`
}
