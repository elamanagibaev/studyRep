package repositories

import (
	"database/sql"
	"log"
	"module3Bit/entities"
)

type ItemRepository interface {
	GetAllItemsRepository() []entities.Item
	GetItemByIDRepository(id int64) entities.Item
	AddItemRepository(item entities.Item)
	UpdateItemRepository(item entities.Item)
	DeleteItemRepository(id int64)
}

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (itemRepository *itemRepository) GetAllItemsRepository() []entities.Item {
	var items []entities.Item
	rows, err := itemRepository.db.Query("select * from items")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var item entities.Item
		rows.Scan(&item.ID, &item.Name, &item.Price, &item.Amount, &item.Promo)
		items = append(items, item)
	}
	return items
}

func (itemRepository *itemRepository) GetItemByIDRepository(id int64) entities.Item {
	var item entities.Item
	row := itemRepository.db.QueryRow("select * from items where id = $1", id)
	err := row.Scan(&item.ID, &item.Name, &item.Price, &item.Amount, &item.Promo)
	if err != nil {
		log.Fatal(err)
		return entities.Item{}
	}
	return item
}

func (itemRepository *itemRepository) AddItemRepository(item entities.Item) {
	_, err := itemRepository.db.Exec("insert into items(name,price,amount, promo) values ($1,$2,$3, $4)", item.Name, item.Price, item.Amount, item.Promo)
	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepository *itemRepository) UpdateItemRepository(item entities.Item) {
	_, err := itemRepository.db.Exec("update items set name = $1, price = $2, amount = $3, promo = $4 where id = $5", item.Name, item.Price, item.Amount, item.Promo, item.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepository *itemRepository) DeleteItemRepository(id int64) {
	_, err := itemRepository.db.Exec("delete from items where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
