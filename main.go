package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"module3Bit/entities"
	"module3Bit/repositories"
	"module3Bit/services"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// Инициализация с БД
	connection := "user=postgres password=Elaman2004123 dbname=postgres  host=localhost port=5432  sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения
	errTwo := db.Ping()
	if errTwo != nil {
		log.Fatal(errTwo)
	}
}

func CloseDB() {
	// Разрыв с БД
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	InitDB()
	defer CloseDB()

	var userRepository repositories.UserRepository
	userRepo := repositories.NewUserRepository(db)
	userRepository = userRepo

	var userService services.UserService
	userServ := services.NewUserService(userRepository)
	userService = userServ

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var user entities.User
			json.NewDecoder(r.Body).Decode(&user)
			userService.AddUserService(user) // обращение через сервисы = handler -> service -> repo

		} else if r.Method == http.MethodGet {
			idUser := r.URL.Query().Get("id")
			if idUser != "" {
				id, err := strconv.Atoi(idUser)
				if err != nil {
					fmt.Println("Null")
				}
				json.NewEncoder(w).Encode(userService.GetUserService(int64(id)))
			} else {
				err := json.NewEncoder(w).Encode(userService.GetAllUsersService())
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
			userService.RemoveUserService(int64(id))

		} else if r.Method == http.MethodPut {
			var user entities.User
			json.NewDecoder(r.Body).Decode(&user)
			userService.UpdateUserService(user)
		}
	})

	var itemRepository repositories.ItemRepository // экземпляр интерфейса
	itemRepo := repositories.NewItemRepository(db) // экземпляр структуры
	itemRepository = itemRepo                      // полимормфизм

	var itemService services.ItemService
	itemServ := services.NewItemService(itemRepository)
	itemService = itemServ

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					return
				}
				json.NewEncoder(w).Encode(itemService.GetItemByIDService(int64(id))) // обращение через сервисы = handler -> service -> repo
			} else {
				json.NewEncoder(w).Encode(itemService.GetAllItemsService()) // обращение через сервисы = handler -> service -> repo
			}
		} else if r.Method == http.MethodPost {
			var item entities.Item
			json.NewDecoder(r.Body).Decode(&item)
			itemService.AddItemService(item) // обращение через сервисы = handler -> service -> repo

		} else if r.Method == http.MethodPut {
			var updItem entities.Item
			json.NewDecoder(r.Body).Decode(&updItem)
			itemService.UpdateItemService(updItem) // обращение через сервисы = handler -> service -> repo

		} else if r.Method == http.MethodDelete {
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Fatal(err)
			}
			itemService.DeleteItemService(int64(id)) // обращение через сервисы = handler -> service -> repo
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
