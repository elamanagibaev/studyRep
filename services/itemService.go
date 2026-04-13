package services

import (
	"module3Bit/dtos"
	"module3Bit/entities"
	"module3Bit/mappers"
	"module3Bit/repositories"

	"github.com/google/uuid"
)

type ItemService interface {
	GetAllItemsService() []dtos.ItemDTO
	GetItemByIDService(id int64) dtos.ItemDTO
	AddItemService(item entities.Item)
	UpdateItemService(item entities.Item)
	DeleteItemService(id int64)
}

type itemService struct {
	itemRepository repositories.ItemRepository
}

func NewItemService(itemRepository repositories.ItemRepository) ItemService {
	return &itemService{itemRepository: itemRepository}
}

func (itemService *itemService) GetAllItemsService() []dtos.ItemDTO {
	return mappers.MapToDTOList(itemService.itemRepository.GetAllItemsRepository())
}

func (itemService *itemService) GetItemByIDService(id int64) dtos.ItemDTO {
	return mappers.MapToDTO(itemService.itemRepository.GetItemByIDRepository(int64(id)))
}

func (itemService *itemService) AddItemService(item entities.Item) {
	if item.Name != "" && item.Amount >= 1 {
		item.Promo = uuid.New().String()
		itemService.itemRepository.AddItemRepository(item)
	}
}

func (itemService *itemService) UpdateItemService(item entities.Item) {
	itemService.itemRepository.UpdateItemRepository(item)
}

func (itemService *itemService) DeleteItemService(id int64) {
	itemService.itemRepository.DeleteItemRepository(int64(id))
}
