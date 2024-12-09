package user

import (
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yordanos-habtamu/EcomGo.git/config"
	"github.com/yordanos-habtamu/EcomGo.git/service/auth"
	"github.com/yordanos-habtamu/EcomGo.git/types"
	"github.com/yordanos-habtamu/EcomGo.git/utils"
)

type Handler struct {
  store types.UserStore
}

func NewHandler(store types.UserStore) *Handler{
	return &Handler{store:store}
}

func(h* Handler) RegisterRoutes(router *mux.Router){
   router.HandleFunc("/login",h.handleLogin).Methods("POST")
   router.HandleFunc("/register",h.handleRegister).Methods("POST")
}
func (h* Handler) handleLogin(w http.ResponseWriter, r*http.Request){
	var payload types.LoginUserPayload
	if err := utils.ParseJson(r,&payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest,err)
	}
	
	if err := utils.Validate.Struct(payload); err!= nil{
		error := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid email or password %s",error))
	}

	u,err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		log.Fatal("user need to register first")
		return
	}
	if !auth.ComparePassword(u.Password,[]byte(payload.Password)){
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid password used"))
		return 
	} 

	secret := []byte(config.Envs.JWT_SECRET)
	token,err := auth.CreateJWT()
	utils.WriteJson(w,http.StatusOK,map[string]string{"token":""})
    


}
func (h* Handler) handleRegister(w http.ResponseWriter, r*http.Request){
// get the payload
var payload types.RegisterUserPayload
if err := utils.ParseJson(r,&payload); err != nil{
	utils.WriteError(w, http.StatusBadRequest,err)
}

if err := utils.Validate.Struct(payload); err!= nil{
	error := err.(validator.ValidationErrors)
	utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid request been sent %s",error))
}
//check if the user exists
 _, err := h.store.GetUserByEmail(payload.Email)
 if err == nil {
  utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user with email %s already exists",payload.Email))
  return
}

//Creating the user if doesn't exist
hashedPassword,err := auth.HashPassword(payload.Password)
if err != nil {
	log.Fatal(err)
}

// The layout for parsing (Go's reference date is "2006-01-02 15:04:05")
layout := "2006-01-02"
dob, err := time.Parse(layout, payload.DoB)
if err != nil {
	log.Fatal(err)
}

err = h.store.CreateUser(types.User{
	FirstName: payload.FirstName,
	LastName: payload.LastName,
	Email: payload.Email,
	Password: hashedPassword,
	DoB: dob,
	Sex : payload.Sex,
})
 
if err != nil {
	utils.WriteError(w, http.StatusInternalServerError,err)
	return
}


}