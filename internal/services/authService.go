package services

import (
	"errors"
	"module3Bit/internal/entities"
	"module3Bit/internal/repositories"
	"module3Bit/pkg/errorsCustom"
)

type AuthService interface {
	AuthUser(user entities.User) error
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (authService *authService) AuthUser(user entities.User) error {
	checkUser, err := authService.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		var notFoundErr errorsCustom.NotFoundError
		if errors.As(err, &notFoundErr) {
			return notFoundErr
		}
		return err
	}

	if checkUser.Password != user.Password {
		return errorsCustom.UnauthorizedError{
			Reason: "Invalid email or password",
		}
	}
	return nil
}
