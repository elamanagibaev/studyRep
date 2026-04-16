package main

import (
	"database/sql"
	"log"
	"module3Bit/handlers"
	"module3Bit/repositories"
	"module3Bit/services"
	"net/http"

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

	var userHandler handlers.UserHandler
	userHandle := handlers.NewUserHandler(userService)
	userHandler = userHandle

	http.HandleFunc("/users", userHandler.HandleRequest)

	var itemRepository repositories.ItemRepository // экземпляр интерфейса
	itemRepo := repositories.NewItemRepository(db) // экземпляр структуры
	itemRepository = itemRepo                      // полимормфизм

	var itemService services.ItemService
	itemServ := services.NewItemService(itemRepository)
	itemService = itemServ

	var itemHandler handlers.ItemHandler
	itemHandle := handlers.NewItemHandler(itemService)
	itemHandler = itemHandle

	http.HandleFunc("/items", itemHandler.HandleRequest)

	server := http.Server{
		Addr: "localhost:4040",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
