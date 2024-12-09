package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error) {
   hash,err := bcrypt.GenerateFromPassword( []byte(password),bcrypt.DefaultCost )
   if err != nil {
	return "" ,err
   }
   return string(hash) ,nil
}

func ComparePassword(password string, plain []byte) bool {
   err := bcrypt.CompareHashAndPassword([]byte(password),plain)
   return err == nil
}