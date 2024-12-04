package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yordanos-habtamu/EcomGo.git/types"
)

type Handler struct {

}

func NewHandler() *Handler{
	return &Handler{}
}

func(h* Handler) RegisterRoutes(router *mux.Router){
   router.HandleFunc("/login",h.handleLogin).Methods("POST")
   router.HandleFunc("/register",h.handleRegister).Methods("POST")
}
func (h* Handler) handleLogin(w http.ResponseWriter, r*http.Request){

}
func (h* Handler) handleRegister(w http.ResponseWriter, r*http.Request){
// get the payload
var payload types.RegisterUserPayload
err := json.NewDecoder(r.Body).Decode(payload)
if err != nil{
	log.Fatal(err)
}

//check if the user exists
}