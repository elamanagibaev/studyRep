package handlers

import (
	"encoding/json"
	"log"
	"module3Bit/entities"
	"module3Bit/services"
	"net/http"
	"strconv"
)

type ItemHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type itemHandler struct {
	itemService services.ItemService
}

func NewItemHandler(itemService services.ItemService) ItemHandler {
	return &itemHandler{itemService: itemService}
}

func (itemHandler *itemHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return
			}
			json.NewEncoder(w).Encode(itemHandler.itemService.GetItemByIDService(int64(id))) // обращение через сервисы = handler -> service -> repo
		} else {
			json.NewEncoder(w).Encode(itemHandler.itemService.GetAllItemsService()) // обращение через сервисы = handler -> service -> repo
		}
	} else if r.Method == http.MethodPost {
		var item entities.Item
		json.NewDecoder(r.Body).Decode(&item)
		itemHandler.itemService.AddItemService(item) // обращение через сервисы = handler -> service -> repo

	} else if r.Method == http.MethodPut {
		var updItem entities.Item
		json.NewDecoder(r.Body).Decode(&updItem)
		itemHandler.itemService.UpdateItemService(updItem) // обращение через сервисы = handler -> service -> repo

	} else if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
		}
		itemHandler.itemService.DeleteItemService(int64(id)) // обращение через сервисы = handler -> service -> repo
	}
}
