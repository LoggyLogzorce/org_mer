package models

type Employee struct {
	Uid        uint8  `json:"uid" gorm:"primary_key"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Patronymic string `json:"patronymic"`
	Address    uint8  `json:"address"`
	Post       uint8  `json:"post"`
	Passport   uint8  `json:"passport"`
	User       uint8  `json:"user"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Photo      string `json:"photo"`
}
