package models

type Customer struct {
	Uid        uint8  `json:"uid" gorm:"primary_key"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Patronymic string `json:"patronymic"`
	Status     uint8  `json:"status"`
	User       uint8  `json:"user"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}
