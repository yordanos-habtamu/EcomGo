package types

import (

	"time"
)

type UserStore interface{
	GetUserByEmail(email string) (*User,error)
	CreateUser(User) error
	GetUserById(id int) (*User,error)
}

type OrderStore interface {
	CreateOrder (Order) (int,error)
	CreateOrderItem(OrderItem) error
}

type ProductStore interface {
	CreateProduct(Product) (error)
	GetProductById(id uint) (*Product, error)
	GetAllProducts() ([]Product, error)
	UpdateProduct(id uint, payload RegisterProductPayload) (*Product, error)
	DeleteProduct(id uint) error
	GetProductsByCategory(category string) ([]Product, error)
	GetProductByName(name string) (*Product, error)
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
type Order struct {
	ID  	uint  		`json:"id"`
    UserID 	uint 	`json:"userId"`
	Total 	float64 	`json:"total"`
	Status 	string 	`json:"status"`
	Address string	`json:"address"`
	CreatedAt string `json:"createdAt"`
	BillingAddress string `json:"billingAddress"`                  
	PaymentMethod  string `json:"paymentMethod"`
	PaymentStatus  string  `json:"paymentStatus"`  
	OrderDate  time.Time   `json:"orderDate"`
	ShipmentDate time.Time `json:"shipmentDate"`              
	DeliveryDate time.Time `json:"deliveryDate"`
    TrackingNumber int `json:"trackingNumber"` 
 }

 type OrderItem struct {
    ID  uint `json:"id"`
	ProductID uint `json:"productId"`
	OrderID uint  `json:"orderId"`
	Quantity uint `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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