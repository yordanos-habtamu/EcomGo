package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/yordanos-habtamu/EcomGo.git/service/cart"
	"github.com/yordanos-habtamu/EcomGo.git/service/product"
	"github.com/yordanos-habtamu/EcomGo.git/service/user"
	"github.com/yordanos-habtamu/EcomGo.git/service/order"

	"github.com/gorilla/mux"
)

type ApiServer struct{
	addr string
	db  *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer{
    return &ApiServer{
    	addr: addr,
    	db:db,
			}		
}

func (s *ApiServer) Run() error{
    
	router := mux.NewRouter();
	subrouter := router.PathPrefix("/api/v1").Subrouter();
    userStore := user.NewStore(s.db)
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore,productStore,userStore)
    cartHandler.RegisterRoutes(subrouter)
  
    log.Printf("Listening on %v",s.addr)
  	 return http.ListenAndServe(s.addr,router)
}