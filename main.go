package main

import (
	"database/sql"
	"log"
	"module3Bit/handlers"
	"module3Bit/repositories"
	"module3Bit/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"

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
	router := mux.NewRouter()

	var userRepository repositories.UserRepository
	userRepo := repositories.NewUserRepository(db)
	userRepository = userRepo

	var userService services.UserService
	userServ := services.NewUserService(userRepository)
	userService = userServ

	var userHandler handlers.UserHandler
	userHandle := handlers.NewUserHandler(userService)
	userHandler = userHandle

	router.HandleFunc("/users", userHandler.HandleRequestGet).Methods("GET")
	router.HandleFunc("/users", userHandler.HandleRequestPost).Methods("POST")
	router.HandleFunc("/users", userHandler.HandleRequestPut).Methods("PUT")
	router.HandleFunc("/users", userHandler.HandleRequestDelete).Methods("DELETE")

	var itemRepository repositories.ItemRepository // экземпляр интерфейса
	itemRepo := repositories.NewItemRepository(db) // экземпляр структуры
	itemRepository = itemRepo                      // полимормфизм

	var itemService services.ItemService
	itemServ := services.NewItemService(itemRepository)
	itemService = itemServ

	var itemHandler handlers.ItemHandler
	itemHandle := handlers.NewItemHandler(itemService)
	itemHandler = itemHandle

	router.HandleFunc("/items", itemHandler.HandleRequestGet).Methods("GET")
	router.HandleFunc("/items", itemHandler.HandleRequestPost).Methods("POST")
	router.HandleFunc("/items", itemHandler.HandleRequestPut).Methods("PUT")
	router.HandleFunc("/items", itemHandler.HandleRequestDelete).Methods("DELETE")

	server := http.Server{
		Addr:    "localhost:4040",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
