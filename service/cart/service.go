package cart

import (
	"fmt"

	"github.com/yordanos-habtamu/EcomGo.git/types"
)

func getCartItemsIDs (items []types.CartItem)([]uint,error){

	productIDs := make([]uint,len(items))
   for i,  item := range items{
	if item.Quantity<=0{
		return nil, fmt.Errorf("Invalid quantity for the product %d",item.ProductID)
	}
	productIDs[i]= item.ProductID
   }
   return productIDs,nil
}

func (h *Handler)CreateOrder(ps []types.Product,items []types.CartItem,userID int)(int,float64,error){
	//check if all the products exist
	productMap := make(map[uint]types.Product)
	for _,product := range ps{
		productMap[product.ID] =product
	}
	// calculate the total product
	if err := checkIfCartIsInStock(items,productMap); err!= nil {
        return 0,0,nil
	}
	totalPrice := calculateTotalPrice(items,productMap)
	//reduce the quantity of the product
	for _,item := range items{
		product := productMap[item.ProductID]
		product.Stock -= item.Quantity
		h.productStore.UpdateProduct(product.ID, types.RegisterProductPayload{
	    	Name: product.Name,
			Description: product.Description,
			Price: product.Price,
			Stock: product.Stock,      
			Catagory: product.Catagory,  
			ImgUrl: product.ImgUrl,
			IsActive: product.IsActive, 
		})
	}
	//create the order
     orderID,err := h.store.CreateOrder(types.Order{
			
		UserID:uint(userID),
		Total:totalPrice,
		Address: 
		BillingAddress:                 
		PaymentMethod 
		PaymentStatus:  
		ShipmentDate time.Time `json:"shipmentDate"`              
		DeliveryDate time.Time `json:"deliveryDate"`
		TrackingNumber int `json:"trackingNumber"` 
	 })
	//create the orderItems

	return 0,totalPrice,nil
}

func checkIfCartIsInStock(cartItems []types.CartItem,products map[uint]types.Product) error {
  if len(cartItems) ==0 {
	return fmt.Errorf("the cart is empty")
  }
  for _,item:= range cartItems {
	product, ok := products[item.ProductID]
	if !ok{
		return fmt.Errorf("product %d is not available in the store,please refresh your cart",product.ID)
	}
	if  product.Stock <item.Quantity {
		return fmt.Errorf("the quantity you are asking is not available %d",product.Stock)
	}
  }
  return nil
}

func calculateTotalPrice(cartItems [] types.CartItem,products map[uint]types.Product) float64 {
	var totalPrice float64;
	for _,item:= range cartItems {
		product := products[item.ProductID]
		totalPrice += float64(product.Price * float64(item.Quantity))
	  }
    return totalPrice
}