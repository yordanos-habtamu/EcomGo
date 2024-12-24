package order

import (
	"database/sql"

	"log"
	"github.com/yordanos-habtamu/EcomGo.git/types"
)

type Store struct {
  db *sql.DB
}
func NewStore (db *sql.DB) *Store {
	return &Store{db:db}
  }

  
func (s *Store) CreateOrder (order types.Order) (int,error){
	res,err := s.db.Exec("INSERT INTO orders (id,user_id,total,status,address,created_at,billing_address,payment_method,order_date,shipment_date,delivery_date,tracking_number) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)",
	order.ID,order.UserID,order.Total,order.Status,order.Address,order.CreatedAt,order.BillingAddress,order.PaymentMethod,order.PaymentStatus,order.OrderDate,order.ShipmentDate,order.DeliveryDate,order.TrackingNumber)
	if err != nil{
	 return 0,err
	}
	
	id,err := res.LastInsertId()
	if err != nil {
		return 0,err
	}
	return int(id),nil
 }
 
 func scanRowsIntoOrders(rows *sql.Rows) (*types.Order,error){
   order := new(types.Order)
   err := rows.Scan(
	  &order.ID,
	  &order.UserID,
	  &order.Total,
	  &order.Status,
	  &order.Address,
	  &order.BillingAddress,                
	  &order.PaymentMethod,  
	  &order.PaymentStatus,   
	  &order.OrderDate,                 
	  &order.DeliveryDate,
	  &order.TrackingNumber, 
   )
   if err != nil{
	 log.Fatal(err)
   }
	 return order , nil
 }
 

 func(s *Store) CreateOrderItem(OrderItem types.OrderItem) error{
   _,err := s.db.Exec("INSERT INTO order_items(id,product_id,order_id,quantity,created_at) VALUES (?,?,?,?,?)",
  OrderItem.ID,OrderItem.ProductID,OrderItem.OrderID,OrderItem.Quantity,OrderItem.CreatedAt)
  if err!= nil {
	return err
  }
  return nil
 }
 
 func scanRowsIntoOrderItems(rows *sql.Rows) (*types.OrderItem,error){
	orderItem := new(types.OrderItem)
	err := rows.Scan(
	   &orderItem.ID,
	   &orderItem.ProductID,
	   &orderItem.OrderID,
	   &orderItem.Quantity,
	   &orderItem.CreatedAt,
	   &orderItem.UpdatedAt, 
	)
	if err != nil{
	  log.Fatal(err)
	}
	  return orderItem , nil
  }