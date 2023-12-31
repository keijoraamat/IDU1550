package main

import (
	"log"
	"net/http"

	"github.com/keijoraamat/IDU1550/ylesanne5/BLL"
	"github.com/keijoraamat/IDU1550/ylesanne5/DAL"
	"github.com/keijoraamat/IDU1550/ylesanne5/PL"
	"github.com/keijoraamat/IDU1550/ylesanne5/initializer"
)

var service PL.ItemServiceInterface

func init() {
	db := initializer.ConnectToDB()
	repo := &DAL.ItemRepository{DB: db}
	service = &BLL.ItemService{Repo: repo}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/items", PL.GetAllItemsHandler(service))
	mux.HandleFunc("/item", PL.CreateItemHandler(service))
	mux.HandleFunc("/item/delete/", PL.DeleteItemHandler(service))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
