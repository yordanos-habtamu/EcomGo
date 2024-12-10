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
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
    var payload types.LoginUserPayload

    // Parse JSON payload
    if err := utils.ParseJson(r, &payload); err != nil {
        log.Printf("Failed to parse JSON: %v", err)
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
        return
    }

    // Validate input
    if err := utils.Validate.Struct(payload); err != nil {
        log.Printf("Validation error: %v", err)
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
        return
    }

    // Check if user exists
    u, err := h.store.GetUserByEmail(payload.Email)
    if err != nil {
        log.Printf("User not found: %v", err)
        utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("user needs to register first"))
        return
    }

    // Verify password
    if !auth.ComparePassword(u.Password, []byte(payload.Password)) {
        log.Printf("Invalid password for user: %s", payload.Email)
        utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
        return
    }

    // Generate JWT
    secret := []byte(config.Envs.JWT_SECRET)
    token, err := auth.CreateJWT(secret, int(u.ID),u.Role)
    if err != nil {
        log.Printf("Error creating JWT: %v", err)
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create token"))
        return
    }

    // Respond with token
    utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})
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
log.Printf("Received DoB: %s", payload.DoB) // Debug print

dob, err := time.Parse(layout, payload.DoB)
if err != nil {
    log.Printf("Error parsing DoB: %v", err)
    http.Error(w, "Invalid Date of Birth format, expected YYYY-MM-DD", http.StatusBadRequest)
    return
}

// Assuming CreateUser is a method of your store that stores the user in the DB
err = h.store.CreateUser(types.User{
    FirstName: payload.FirstName,
    LastName:  payload.LastName,
    Email:     payload.Email,
    Password:  hashedPassword,
    DoB:       dob,
    Sex:       payload.Sex,
    Role:      payload.Role,
})
if err != nil {
    log.Printf("Error creating user: %v", err)
    http.Error(w, "Error creating user", http.StatusInternalServerError)
    return
}
if err != nil {
	utils.WriteError(w, http.StatusInternalServerError,err)
	return
}


}