package types

import "time"

type UserStore interface{
	GetUserByEmail(email string) (*User,error)
	CreateUser(User) error
	GetUserById(id int) (*User,error)

}
type User struct {
	ID  uint  `json:"Id"`
	FirstName string `json:"firstName"`
	LastName string   `json:"lastName"`
	Sex string `json:"sex"`
	Email string     `json:"email"`
	DoB time.Time   `json:"DoB"`
	Password string  `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}


type RegisterUserPayload struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string   `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"required,email"` 
	DoB       string    `json:"DoB" validate:"required"`
	Sex       string      `json:"sex" validate:"required"`
	Password   string      `json:"password" validate:"required,min=6,max=12"`
}
