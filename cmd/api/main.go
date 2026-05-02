package main

import (
	"database/sql"
	"log"
	handlers2 "module3Bit/internal/handlers"
	repositories2 "module3Bit/internal/repositories"
	services2 "module3Bit/internal/services"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// Инициализация с БД
	var err error
	connection := "user=postgres password=Elaman2004123 dbname=postgres  host=localhost port=5432  sslmode=disable"
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal(errPing)
	}
}

func CloseDB() {
	// Разрыв с БД
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: ", r.URL.Path, "METHOD: ", r.Method)
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println("delta time: ", time.Since(start))
		log.Println("Response to client has been sent")
	})
}

func RecoveryDBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println("PANIC", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	InitDB()
	defer CloseDB()
	router := mux.NewRouter()

	var userRepository repositories2.UserRepository
	userRepo := repositories2.NewUserRepository(db)
	userRepository = userRepo

	var userService services2.UserService
	userServ := services2.NewUserService(userRepository)
	userService = userServ

	var userHandler handlers2.UserHandler
	userHandle := handlers2.NewUserHandler(userService)
	userHandler = userHandle

	var authService services2.AuthService
	authServ := services2.NewAuthService(userRepository)
	authService = authServ

	var authHandler handlers2.AuthHandler
	authHand := handlers2.NewAuthHandler(authService)
	authHandler = authHand

	router.HandleFunc("/users", userHandler.HandleRequestGet).Methods("GET")
	router.HandleFunc("/users", userHandler.HandleRequestPost).Methods("POST")
	router.HandleFunc("/users", userHandler.HandleRequestPut).Methods("PUT")
	router.HandleFunc("/users", userHandler.HandleRequestDelete).Methods("DELETE")
	router.HandleFunc("/auth", authHandler.BasicAuth).Methods("GET")

	var itemRepository repositories2.ItemRepository // экземпляр интерфейса
	itemRepo := repositories2.NewItemRepository(db) // экземпляр структуры
	itemRepository = itemRepo                       // полимормфизм

	var itemService services2.ItemService
	itemServ := services2.NewItemService(itemRepository)
	itemService = itemServ

	var itemHandler handlers2.ItemHandler
	itemHandle := handlers2.NewItemHandler(itemService)
	itemHandler = itemHandle

	router.HandleFunc("/items", itemHandler.HandleRequestGet).Methods("GET")
	router.HandleFunc("/items", itemHandler.HandleRequestPost).Methods("POST")
	router.HandleFunc("/items", itemHandler.HandleRequestPut).Methods("PUT")
	router.HandleFunc("/items", itemHandler.HandleRequestDelete).Methods("DELETE")

	middleware := RecoveryDBMiddleware(router) // сначала обработаем соединение с БД
	middleware = LoggingMiddleware(middleware) // после перезаписывание при логирование

	server := http.Server{
		Addr:    "localhost:4040",
		Handler: middleware,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
