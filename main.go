package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"module3Bit/entities"
	"module3Bit/mappers"
	"module3Bit/repositories"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var user entities.User
			json.NewDecoder(r.Body).Decode(&user)
			userRepository.AddUser(user)
		} else if r.Method == http.MethodGet {
			idUser := r.URL.Query().Get("id")
			if idUser != "" {
				id, err := strconv.Atoi(idUser)
				if err != nil {
					fmt.Println("Null")
				}
				json.NewEncoder(w).Encode(mappers.MapToUserDTO(userRepository.GetUser(int64(id))))
			} else {
				err := json.NewEncoder(w).Encode(mappers.MapToUserDTOList(userRepository.GetAllUsers()))
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
			userRepository.RemoveUser(int64(id))
		} else if r.Method == http.MethodPut {
			var user entities.User
			json.NewDecoder(r.Body).Decode(&user)
			userRepository.UpdateUser(user)
		}
	})

	var itemRepository repositories.ItemRepository // экземпляр интерфейса
	itemRepo := repositories.NewItemRepository(db) // экземпляр структуры
	itemRepository = itemRepo                      // полимормфизм

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					return
				}
				json.NewEncoder(w).Encode(mappers.MapToDTO(itemRepository.GetItemByID(int64(id))))
			} else {
				json.NewEncoder(w).Encode(mappers.MapToDTOList(itemRepository.GetAllItems()))
			}
		} else if r.Method == http.MethodPost {
			var item entities.Item
			json.NewDecoder(r.Body).Decode(&item)
			item.Promo = uuid.New().String()
			itemRepository.AddItem(item)
		} else if r.Method == http.MethodPut {
			var updItem entities.Item
			json.NewDecoder(r.Body).Decode(&updItem)
			itemRepository.UpdateItem(updItem)
		} else if r.Method == http.MethodDelete {
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Fatal(err)
			}
			itemRepository.DeleteItem(int64(id))
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
