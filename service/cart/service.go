package cart

import (
    "fmt"
    "time"

    "github.com/yordanos-habtamu/EcomGo.git/types"
)

func getCartItemsIDs(items []types.CartItem) ([]uint, error) {
    productIDs := make([]uint, len(items))
    for i, item := range items {
        if item.Quantity <= 0 {
            return nil, fmt.Errorf("Invalid quantity for the product %d", item.ProductID)
        }
        productIDs[i] = item.ProductID
    }
    return productIDs, nil
}

func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {
    // Check if all the products exist
    productMap := make(map[uint]types.Product)
    for _, product := range ps {
        productMap[product.ID] = product
    }

    // Calculate the total product
    if err := checkIfCartIsInStock(items, productMap); err != nil {
        return 0, 0, err
    }
    totalPrice := calculateTotalPrice(items, productMap)

    // Reduce the quantity of the product
    for _, item := range items {
        product := productMap[item.ProductID]
        product.Stock -= item.Quantity
        h.productStore.UpdateProduct(product.ID, types.RegisterProductPayload{
            Name:        product.Name,
            Description: product.Description,
            Price:       product.Price,
            Stock:       product.Stock,
            Catagory:    product.Catagory,
            ImgUrl:      product.ImgUrl,
            IsActive:    product.IsActive,
        })
    }

    // Create the order
    orderID, err := h.store.CreateOrder(types.Order{
        UserID:         uint(userID),
        Total:          totalPrice,
        Address:        "Shipping Address", // Replace with actual address
        BillingAddress: "Billing Address",  // Replace with actual billing address
        PaymentMethod:  "Payment Method",   // Replace with actual payment method
        PaymentStatus:  "Payment Status",   // Replace with actual payment status
        ShipmentDate:   time.Time{},        // Replace with actual shipment date
        DeliveryDate:   time.Time{},        // Replace with actual delivery date
        TrackingNumber: 0,                  // Replace with actual tracking number
    })
    if err != nil {
        return 0, 0, err
    }

    // Create the order items
    for _, item := range items {
        err := h.store.CreateOrderItem(types.OrderItem{
            OrderID:   uint(orderID),
            ProductID: item.ProductID,
            Quantity:  uint(item.Quantity),
        })
        if err != nil {
            return 0, 0, err
        }
    }

    return orderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[uint]types.Product) error {
    if len(cartItems) == 0 {
        return fmt.Errorf("the cart is empty")
    }
    for _, item := range cartItems {
        product, ok := products[item.ProductID]
        if !ok {
            return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
        }
        if product.Stock < item.Quantity {
            return fmt.Errorf("the quantity you are asking is not available %d", product.Stock)
        }
    }
    return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[uint]types.Product) float64 {
    var totalPrice float64
    for _, item := range cartItems {
        product := products[item.ProductID]
        totalPrice += float64(product.Price * float64(item.Quantity))
    }
    return totalPrice
}