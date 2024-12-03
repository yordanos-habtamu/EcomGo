package main

import (
	"log"

	"github.com/yordanos-habtamu/EcomGo.git/cmd/api"
)

func main() {

	server := api.NewApiServer(":8080",nil)
   if error:= server.Run(); error!=nil{
	log.Fatal(error)
   }
}