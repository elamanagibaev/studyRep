package handlers

import (
	"encoding/json"
	"errors"
	"module3Bit/internal/entities"
	"module3Bit/internal/services"
	"module3Bit/pkg/errorsCustom"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type AuthHandler interface {
	BasicAuth(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (h authHandler) BasicAuth(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid Auth header"})
		return
	}
	var user entities.User
	user.Email = email
	user.Password = password
	err := h.authService.AuthUser(user)
	if err != nil {
		var unAuthErr errorsCustom.UnauthorizedError
		if errors.As(err, &unAuthErr) {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": unAuthErr.Error()})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "Success Auth"})
}
