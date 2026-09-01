package main

import (
	"fmt"
	"net/http"
	"inventory/internal/handler"
)

func main(){
	fmt.Println("Starting Application .....")
	fmt.Println("-----------WELCOME TO INVENTORY MANAGMENT APPPLICATION-----------")
	
	mux :=  http.NewServeMux()
	handler.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)

	fmt.Println("Stopping Application .....")
	fmt.Println("-----------BYE-----------")
}