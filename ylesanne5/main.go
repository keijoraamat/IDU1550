package main

import (
	"log"
	"net/http"

	"github.com/keijoraamat/IDU1550/ylesanne5/PL"
	"github.com/keijoraamat/IDU1550/ylesanne5/initializer"
)

func init() {
	initializer.ConnectToDB()
}

func main() {
	http.HandleFunc("/items", PL.GetAllItemsHandler)
	http.HandleFunc("/item", PL.CreateItemHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
