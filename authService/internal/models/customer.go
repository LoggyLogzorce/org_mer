package models

type Customer struct {
	Uid        uint8  `json:"uid" gorm:"primary_key"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Otchestvo  string `json:"otchestvo"`
	Status     uint8  `json:"status"`
	Polzovatel uint8  `json:"polzovatel"`
	Telephone  string `json:"telephone"`
	Email      string `json:"email"`
}

func (_ *Customer) TableName() string {
	return "polzovateli"
}
