package main

import (
	"log"
	"net/http"

	"github.com/keijoraamat/IDU1550/ylesanne5/BLL"
	"github.com/keijoraamat/IDU1550/ylesanne5/DAL"
	"github.com/keijoraamat/IDU1550/ylesanne5/PL"
	"github.com/keijoraamat/IDU1550/ylesanne5/initializer"
)

var service PL.ItemService

func init() {
	db := initializer.ConnectToDB()
	repo := &DAL.ItemRepository{DB: db}
	service = &BLL.ItemService{Repo: repo}
}

func main() {
	http.HandleFunc("/items", PL.GetAllItemsHandler)
	http.HandleFunc("/item", PL.CreateItemHandler(service))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
