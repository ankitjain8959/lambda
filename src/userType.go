package src

type User struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Age         float64 `json:"age"`
	Address     string  `json:"address"`
	Phone       string  `json:"phone"`
	DateOfBirth string  `json:"dateOfBirth"`
}