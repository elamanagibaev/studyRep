package mappers

import (
	"module3Bit/dtos"
	"module3Bit/entities"
)

func MapToDTO(item entities.Item) dtos.ItemDTO {
	var itemDTO dtos.ItemDTO
	itemDTO.ID = item.ID
	itemDTO.Name = item.Name
	itemDTO.Price = item.Price
	itemDTO.Amount = item.Amount
	return itemDTO
}

func MapToDTOList(items []entities.Item) []dtos.ItemDTO {
	var itemsDTOs []dtos.ItemDTO
	for i := 0; i < len(items); i++ {
		var itemDTO dtos.ItemDTO
		itemDTO = MapToDTO(items[i])
		itemsDTOs = append(itemsDTOs, itemDTO)
	}
	return itemsDTOs
}

func MapToUserDTO(user entities.User) dtos.UserDTO {
	var userDTO dtos.UserDTO
	userDTO.ID = user.ID
	userDTO.Email = user.Email
	return userDTO
}

func MapToUserDTOList(users []entities.User) []dtos.UserDTO {
	var userDtoList []dtos.UserDTO
	for i := 0; i < len(users); i++ {
		var userDTO dtos.UserDTO
		userDTO = MapToUserDTO(users[i])
		userDtoList = append(userDtoList, userDTO)
	}
	return userDtoList
}
