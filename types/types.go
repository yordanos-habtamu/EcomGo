package types

import "time"

type RegisterUserPayload struct {
	FirstName string  `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string    `json:"email"`
	DoB       time.Time    `json:"age"`
	Sex       string      `json:"sex"`
	Password   string      `json:"password"`
}