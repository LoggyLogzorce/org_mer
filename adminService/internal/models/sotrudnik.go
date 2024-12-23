package models

type Sotrudnik struct {
	IdSotrudnika   uint8  `json:"id_sotrudnika"`
	Familiya       string `json:"familiya"`
	Imya           string `json:"imya"`
	Otchestvo      string `json:"otchestvo"`
	Address        string `json:"address"`
	Doljnost       string `json:"doljnost"`
	Passport       string `json:"passport"`
	IdPolzovatelya uint8  `json:"id_polzovatelya"`
	Telephone      string `json:"telephone"`
	Email          string `json:"email"`
	Photo          string `json:"photo"`
}

func (_ *Sotrudnik) TableName() string {
	return "sotrudniki"
}
