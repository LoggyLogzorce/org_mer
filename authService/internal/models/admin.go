package models

type Sotrudnik struct {
	Uid        uint8  `json:"uid" gorm:"primary_key"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Otchestvo  string `json:"otchestvo"`
	Address    uint8  `json:"address"`
	Doljnost   uint8  `json:"doljnost"`
	Passport   uint8  `json:"passport"`
	Polzovatel uint8  `json:"user"`
	Telephone  string `json:"telephone"`
	Email      string `json:"email"`
	Photo      string `json:"photo"`
}

func (_ *Sotrudnik) TableName() string {
	return "sotrudniki"
}
