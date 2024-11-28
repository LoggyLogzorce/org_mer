package models

type User struct {
	IdPolzovatelya   uint8  `json:"id_polzovatelya" gorm:"primary_key"`
	Login            string `json:"login"`
	Password         string `json:"passport"`
	NaimenovanieRoli string `json:"role"`
}

func (_ *User) TableName() string {
	return "polzovateli"
}
