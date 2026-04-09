package repositories

import (
	"database/sql"
	"log"
	"module3Bit/entities"
)

type ItemRepository interface {
	GetAllItems() []entities.Item
	GetItemByID(id int64) entities.Item
	AddItem(item entities.Item)
	UpdateItem(item entities.Item)
	DeleteItem(id int64)
}

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	// норм версия для понимания
	//var itemRepo itemRepository
	//itemRepo.db = db
	//return &itemRepo
	
	return &itemRepository{db: db}
}

func (itemRepository *itemRepository) GetAllItems() []entities.Item {
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

func (itemRepository *itemRepository) GetItemByID(id int64) entities.Item {
	var item entities.Item
	row := itemRepository.db.QueryRow("select * from items where id = $1", id)
	err := row.Scan(&item.ID, &item.Name, &item.Price, &item.Amount, &item.Promo)
	if err != nil {
		log.Fatal(err)
		return entities.Item{}
	}
	return item
}

func (itemRepository *itemRepository) AddItem(item entities.Item) {
	_, err := itemRepository.db.Exec("insert into items(name,price,amount, promo) values ($1,$2,$3, $4)", item.Name, item.Price, item.Amount, item.Promo)
	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepository *itemRepository) UpdateItem(item entities.Item) {
	_, err := itemRepository.db.Exec("update items set name = $1, price = $2, amount = $3, promo = $4 where id = $5", item.Name, item.Price, item.Amount, item.Promo, item.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepository *itemRepository) DeleteItem(id int64) {
	_, err := itemRepository.db.Exec("delete from items where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
