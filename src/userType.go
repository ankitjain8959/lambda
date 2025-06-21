package src

type User struct {
	Id          string  `json:"id" bson:"_id"` // Use json tag for JSON serialization and bson tag for MongoDB compatibility
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Age         float64 `json:"age"`
	Address     string  `json:"address"`
	Phone       string  `json:"phone"`
	DateOfBirth string  `json:"dateOfBirth" bson:"dateOfBirth"`
	AtType      string  `json:"@type" bson:"atType"`
}
