package main

import (
	"encoding/json"
	"log"
	"module3Bit/db"
	"module3Bit/entities"
	"module3Bit/mappers"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var user entities.User
			json.NewDecoder(r.Body).Decode(&user)
			db.AddUser(user)
		} else if r.Method == http.MethodGet {
			idUser := r.URL.Query().Get("id")
			if idUser != "" {
				id, err := strconv.Atoi(idUser)
				if err != nil {
					log.Fatal(err)
				}
				json.NewEncoder(w).Encode(mappers.MapToUserDTO(db.GetUser(int64(id))))
			} else {
				err := json.NewEncoder(w).Encode(mappers.MapToUserDTOList(db.GetAllUsers()))
				if err != nil {
					log.Fatal(err)
				}
			}

		} else if r.Method == http.MethodDelete {
			idUser := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idUser)
			if err != nil {
				log.Fatal(err)
			}
			db.RemoveUser(int64(id))

		} else if r.Method == http.MethodPut {
			var user entities.User
			json.NewDecoder(r.Body).Decode(&user)
			db.UpdateUser(user)
		}
	})

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					return
				}
				json.NewEncoder(w).Encode(mappers.MapToDTO(db.GetItems(int64(id))))
			} else {
				json.NewEncoder(w).Encode(mappers.MapToDTOList(db.GetAllItems()))
			}
		} else if r.Method == http.MethodPost {
			var item entities.Item
			json.NewDecoder(r.Body).Decode(&item)
			item.Promo = uuid.New().String()
			db.AddItem(item)
		} else if r.Method == http.MethodPut {
			var updItem entities.Item
			json.NewDecoder(r.Body).Decode(&updItem)
			db.UpdateItem(updItem)
		} else if r.Method == http.MethodDelete {
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Fatal(err)
			}
			db.DeleteItem(int64(id))
		}
	})

	server := http.Server{
		Addr: "localhost:4040",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
