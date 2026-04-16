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
	HandleRequestGet(w http.ResponseWriter, r *http.Request)
	HandleRequestPost(w http.ResponseWriter, r *http.Request)
	HandleRequestPut(w http.ResponseWriter, r *http.Request)
	HandleRequestDelete(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (userHandler *userHandler) HandleRequestGet(w http.ResponseWriter, r *http.Request) {
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
}

func (userHandler *userHandler) HandleRequestPost(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	userHandler.userService.AddUserService(user)
}

func (userHandler *userHandler) HandleRequestPut(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	userHandler.userService.UpdateUserService(user)
}

func (userHandler *userHandler) HandleRequestDelete(w http.ResponseWriter, r *http.Request) {
	idUser := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		log.Fatal(err)
	}
	userHandler.userService.RemoveUserService(int64(id))
}
