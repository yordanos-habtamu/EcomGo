package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yordanos-habtamu/EcomGo.git/config"
)

func CreateJWT(secret []byte,userID int , role string ) (string,error){
 expiration  := time.Second * time.Duration (config.Envs.JWTExpiration)

 token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
	"userID" : strconv.Itoa(userID),
    "role":role,
	"expiredAt":time.Now().Add(expiration).Unix(),
 })
 tokenString,err := token.SignedString(secret)
 if err != nil {
	return "" ,err
 }
 return tokenString,nil
}

