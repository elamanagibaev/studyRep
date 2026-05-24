package main

import (
	"database/sql"
	"log"
	handlers2 "module3Bit/internal/handlers"
	repositories2 "module3Bit/internal/repositories"
	services2 "module3Bit/internal/services"
	"module3Bit/internal/utils"
	"net/http"
	"strings"
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

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtHeader := r.Header.Get("Authorization")
		if jwtHeader != "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(jwtHeader, "Bearer ")
		token, err := utils.VerifyJwtToken(tokenStr)
		if err != nil || !token.Valid {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	InitDB()
	defer CloseDB()

	router := mux.NewRouter()
	itemRouter := router.PathPrefix("/items").Subrouter()
	itemRouter.Use(JWTMiddleware)

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

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.BasicAuth).Methods("POST")
	authRouter.HandleFunc("/register", authHandler.Register).Methods("POST")

	var itemRepository repositories2.ItemRepository // экземпляр интерфейса
	itemRepo := repositories2.NewItemRepository(db) // экземпляр структуры
	itemRepository = itemRepo                       // полимормфизм

	var itemService services2.ItemService
	itemServ := services2.NewItemService(itemRepository)
	itemService = itemServ

	var itemHandler handlers2.ItemHandler
	itemHandle := handlers2.NewItemHandler(itemService)
	itemHandler = itemHandle

	itemRouter.HandleFunc("", itemHandler.HandleRequestGet).Methods("GET")
	itemRouter.HandleFunc("", itemHandler.HandleRequestPost).Methods("POST")
	itemRouter.HandleFunc("", itemHandler.HandleRequestPut).Methods("PUT")
	itemRouter.HandleFunc("", itemHandler.HandleRequestDelete).Methods("DELETE")

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
