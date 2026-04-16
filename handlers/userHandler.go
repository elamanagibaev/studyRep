package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"module3Bit/entities"
	"module3Bit/services"
	"net/http"
	"strconv"
)

type UserHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (userHandler *userHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user entities.User
		json.NewDecoder(r.Body).Decode(&user)
		userHandler.userService.AddUserService(user) // обращение через сервисы = handler -> service -> repo

	} else if r.Method == http.MethodGet {
		idUser := r.URL.Query().Get("id")
		if idUser != "" {
			id, err := strconv.Atoi(idUser)
			if err != nil {
				fmt.Println("Null")
			}
			json.NewEncoder(w).Encode(userHandler.userService.GetUserService(int64(id)))
		} else {
			err := json.NewEncoder(w).Encode(userHandler.userService.GetAllUsersService())
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
		userHandler.userService.RemoveUserService(int64(id))

	} else if r.Method == http.MethodPut {
		var user entities.User
		json.NewDecoder(r.Body).Decode(&user)
		userHandler.userService.UpdateUserService(user)
	}
}
