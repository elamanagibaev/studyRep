package dtos

type ItemDTO struct {
	ID     int64
	Name   string
	Price  float64
	Amount int
}

type UserDTO struct {
	ID    int64
	Email string
}
