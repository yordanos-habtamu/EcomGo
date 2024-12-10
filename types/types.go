package types

import (

	"time"
)

type UserStore interface{
	GetUserByEmail(email string) (*User,error)
	CreateUser(User) error
	GetUserById(id int) (*User,error)
}


type ProductStore interface {
	CreateProduct(payload RegisterProductPayload) (Product, error)
	GetProductByID(id uint) (Product, error)
	GetAllProducts() ([]Product, error)
	UpdateProduct(id uint, payload RegisterProductPayload) (Product, error)
	DeleteProduct(id uint) error
	GetProductsByCategory(category string) ([]Product, error)
	GetProductByName(name string) (Product, error)
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
	Role   string  `json:"role"`
}

type Product struct {
	ID uint `json:"Id"`
	Name string `json:"name"`
	Description string `json:"description"`
    Price float64 `json:"price"`
	Stock int `json:"stock"`
	Catagory string `json:"catagory"`
	ImgUrl string `json:"imgUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive bool `json:"isActive"`
}

type RegisterProductPayload struct {
	Name string `json:"name"`
	Description string `json:"description"`
    Price float64 `json:"price"`
	Stock int `json:"stock"`
	Catagory string `json:"catagory"`
	ImgUrl string `json:"imgUrl"`
	IsActive bool `json:"isActive"`
}

type RegisterUserPayload struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string   `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"required,email"` 
	DoB       string    `json:"DoB" validate:"required"`
	Sex       string      `json:"sex" validate:"required"`
	Password   string      `json:"password" validate:"required,min=6,max=12"`
	Role       string      `json:"role"`
}

type LoginUserPayload struct {
	Email string 	`json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}