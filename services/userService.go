package services

import (
	"module3Bit/dtos"
	"module3Bit/entities"
	"module3Bit/mappers"
	"module3Bit/repositories"
)

type UserService interface {
	AddUserService(user entities.User)
	GetUserService(id int64) dtos.UserDTO
	GetAllUsersService() []dtos.UserDTO
	RemoveUserService(id int64)
	UpdateUserService(user entities.User)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (userService *userService) AddUserService(user entities.User) {
	userService.userRepository.AddUser(user)
}

func (userService *userService) GetUserService(id int64) dtos.UserDTO {
	return mappers.MapToUserDTO(userService.userRepository.GetUser(int64(id)))
}

func (userService *userService) GetAllUsersService() []dtos.UserDTO {
	return mappers.MapToUserDTOList(userService.userRepository.GetAllUsers())
}

func (userService *userService) RemoveUserService(id int64) {
	userService.userRepository.RemoveUser(int64(id))
}

func (userService *userService) UpdateUserService(user entities.User) {
	userService.userRepository.UpdateUser(user)
}
