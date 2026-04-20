package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"module3Bit/internal/entities"
)

type UserRepository interface {
	AddUser(user entities.User)
	GetUser(id int64) entities.User
	GetAllUsers() []entities.User
	RemoveUser(id int64)
	UpdateUser(user entities.User)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) AddUser(user entities.User) {
	_, err := userRepository.db.Exec("insert into users(email, password) values ($1, $2)", user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func (userRepository *userRepository) GetUser(id int64) entities.User {
	var user entities.User
	rows := userRepository.db.QueryRow("select * from users where id = $1", id)
	errFindUser := rows.Scan(&user.ID, &user.Email, &user.Password)
	if errFindUser != nil {
		return entities.User{}
	}
	return user
}

func (userRepository *userRepository) GetAllUsers() []entities.User {
	var users []entities.User
	rows, err := userRepository.db.Query("select * from users")
	if err != nil {
		fmt.Println("Ошибка", err)
	}

	for rows.Next() {
		var user entities.User
		rows.Scan(&user.ID, &user.Email, &user.Password)
		users = append(users, user)
	}
	return users
}

func (userRepository *userRepository) RemoveUser(id int64) {
	_, err := userRepository.db.Exec("delete from users where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func (userRepository *userRepository) UpdateUser(user entities.User) {
	_, err := userRepository.db.Exec("update users set email = $1, password = $2 where id = $3", user.Email, user.Password, user.ID)
	if err != nil {
		log.Fatal(err)
	}
}
