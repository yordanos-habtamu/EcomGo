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
	Age  uint     `json:"age"`
	Sex string `json:"sex"`
	Email string     `json:"email"`
	DoB time.Time     `json:"DoB"`
	Password string  `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}


type RegisterUserPayload struct {
	FirstName string  `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string    `json:"email"`
	Age       uint        `json:"age"`
	DoB       time.Time    `json:"DoB"`
	Sex       string      `json:"sex"`
	Password   string      `json:"password"`
}
