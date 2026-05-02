package handlers

import (
	"encoding/json"
	"module3Bit/internal/entities"
	"module3Bit/internal/services"
	"net/http"
	"strconv"
)

type ItemHandler interface {
	HandleRequestGet(w http.ResponseWriter, r *http.Request)
	HandleRequestPost(w http.ResponseWriter, r *http.Request)
	HandleRequestPut(w http.ResponseWriter, r *http.Request)
	HandleRequestDelete(w http.ResponseWriter, r *http.Request)
}

type itemHandler struct {
	itemService services.ItemService
}

func NewItemHandler(itemService services.ItemService) ItemHandler {
	return &itemHandler{itemService: itemService}
}

func (itemHandler *itemHandler) HandleRequestGet(w http.ResponseWriter, r *http.Request) {
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
}

func (itemHandler *itemHandler) HandleRequestPost(w http.ResponseWriter, r *http.Request) {
	var item entities.Item
	json.NewDecoder(r.Body).Decode(&item)
	itemHandler.itemService.AddItemService(item)
}

func (itemHandler *itemHandler) HandleRequestPut(w http.ResponseWriter, r *http.Request) {
	var updItem entities.Item
	json.NewDecoder(r.Body).Decode(&updItem)
	itemHandler.itemService.UpdateItemService(updItem)
}

func (itemHandler *itemHandler) HandleRequestDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	itemHandler.itemService.DeleteItemService(int64(id)) // обращение через сервисы = handler -> service -> repo
}
