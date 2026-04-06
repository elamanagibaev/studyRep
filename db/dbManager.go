package db

import (
	"database/sql"
	"log"
	"module3Bit/entities"
)

var db *sql.DB

func InitDB() {
	// Инициализация с БД
	connection := "user=postgres password=Elaman2004123 dbname=postgres  host=localhost port=5432  sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения
	errTwo := db.Ping()
	if errTwo != nil {
		log.Fatal(errTwo)
	}
}

func CloseDB() {
	// Разрыв с БД
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllItems() []entities.Item {
	var items []entities.Item
	rows, err := db.Query("select * from items")
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

func GetItems(id int64) entities.Item {
	var item entities.Item
	row := db.QueryRow("select * from items where id = $1", id)
	err := row.Scan(&item.ID, &item.Name, &item.Price, &item.Amount, &item.Promo)
	if err != nil {
		log.Fatal(err)
		return entities.Item{}
	}
	return item
}

func AddItem(item entities.Item) {
	_, err := db.Exec("insert into items(name,price,amount, promo) values ($1,$2,$3, $4)", item.Name, item.Price, item.Amount, item.Promo)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateItem(item entities.Item) {
	_, err := db.Exec("update items set name = $1, price = $2, amount = $3, promo = $4 where id = $5", item.Name, item.Price, item.Amount, item.Promo, item.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteItem(id int64) {
	_, err := db.Exec("delete from items where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func AddUser(user entities.User) {
	_, err := db.Exec("insert into users(email, password) values ($1, $2)", user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser() []entities.User {
	var users []entities.User
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user entities.User
		rows.Scan(&user.ID, &user.Email, &user.Password)
		users = append(users, user)
	}
	return users
}
