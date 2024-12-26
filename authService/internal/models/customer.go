package models

type Zakazchik struct {
	IdZakazchika        uint8  `json:"id_zakazchika" gorm:"primaryKey;column:id_zakazchika"`
	Familiya            string `json:"familiya" gorm:"column:familiya"`
	Imya                string `json:"imya" gorm:"column:imya"`
	Otchestvo           string `json:"otchestvo" gorm:"column:otchestvo"`
	IdStatusaZakazchika uint8  `json:"statusZakazchika" gorm:"column:id_statusa_zakazchika"`
	IdPolzovatelya      uint8  `json:"polzovatelya" gorm:"column:id_polzovatelya"`
	Telephone           string `json:"telephone" gorm:"column:telephone"`
	Email               string `json:"email" gorm:"column:email"`
}

func (_ *Zakazchik) TableName() string {
	return "zakazchiki"
}
