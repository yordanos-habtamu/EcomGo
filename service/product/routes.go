package product

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yordanos-habtamu/EcomGo.git/service/middleware"
	"github.com/yordanos-habtamu/EcomGo.git/types"
	"github.com/yordanos-habtamu/EcomGo.git/utils"
)


type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler{
	return &Handler{store:store}
}

// RegisterRoutes registers the routes for product-related operations
func (h *Handler) RegisterRoutes(router *mux.Router) {
	productMiddleware:= middleware.JwtMiddleware("admin","seller")
	adminOnlyMiddleware:= middleware.JwtMiddleware("admin")
	router.HandleFunc("/products",  productMiddleware(h.handleCreateProduct)).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", h.handleGetProductByID).Methods("GET")
	router.HandleFunc("/products", h.handleGetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", adminOnlyMiddleware(h.handleUpdateProduct)).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", adminOnlyMiddleware( h.handleDeleteProduct)).Methods("DELETE")
	router.HandleFunc("/products/category/{category}", h.handleGetProductsByCategory).Methods("GET")
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
   err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Catagory:    payload.Catagory,
		ImgUrl:      payload.ImgUrl,
	})
	if err != nil {
		log.Printf("Error creating product : %v",err)
		http.Error(w,"Error creating user",http.StatusInternalServerError)
	    return
		} 

		utils.WriteJson(w,http.StatusCreated,map[string]string{"message":"product created successfully"})


}

func (h *Handler) handleGetProductByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path parameters
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product ID is missing in the URL"))
		return
	}

	// Convert the ID to the appropriate type (e.g., uint)
	productID, err := utils.StringToUint(id) // Assuming you have a utility to convert string to uint
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID format"))
		return
	}

	// Fetch the product by ID from the database
	product, err := h.store.GetProductById(productID)
	if err != nil {
		log.Printf("Failed to fetch product: %v", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	// Return the product as a JSON response
	utils.WriteJson(w, http.StatusOK, product)
}


func (h *Handler) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch all products from the store
	products, err := h.store.GetAllProducts()
	if err != nil {
		log.Printf("Failed to fetch products: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch products"))
		return
	}

	// Return the products as a JSON response
	utils.WriteJson(w, http.StatusOK, products)
}


func (h *Handler) handleUpdateProduct(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product ID is missing in the URL"))
		return
	}
	productID, err := utils.StringToUint(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID format"))
		return
	} //
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
    
	p,err:= h.store.UpdateProduct(productID,payload)
	if err != nil {
		log.Printf("Failed to update the product: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch products"))
		return
	}
	utils.WriteJson(w, http.StatusOK, p)
}

func (h *Handler) handleDeleteProduct(w http.ResponseWriter,r *http.Request){
    vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product ID is missing in the URL"))
		return
	}
	productID, err := utils.StringToUint(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID format"))
		return
	} //
    
	k:= h.store.DeleteProduct(productID)
	if k!=nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to delete the product %v",k))
		return
	}
  
	utils.WriteJson(w, http.StatusOK,map[string]string{"message":"Removed the product succefully"})

}

func (h *Handler) handleGetProductsByCategory(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}


}