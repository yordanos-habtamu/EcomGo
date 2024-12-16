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