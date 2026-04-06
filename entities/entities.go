package entities

type Item struct {
	ID     int64
	Name   string
	Price  float64
	Amount int
	Promo  string
}

type User struct {
	ID       int64
	Email    string
	Password string
}
