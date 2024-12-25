package models

type Zakazchik struct {
	IdZakazchika        uint8            `json:"id_zakazchika" gorm:"primaryKey;column:id_zakazchika"`
	Familiya            string           `json:"familiya" gorm:"column:familiya"`
	Imya                string           `json:"imya" gorm:"column:imya"`
	Otchestvo           string           `json:"otchestvo" gorm:"column:otchestvo"`
	IdStatusaZakazchika uint8            `json:"statusZakazchika" gorm:"column:id_statusa_zakazchika"`
	Telephone           string           `json:"telephone" gorm:"column:telephone"`
	Email               string           `json:"email" gorm:"column:email"`
	StatusZakazchika    StatusZakazchika `gorm:"foreignKey:IdStatusaZakazchika;references:IdStatusa"`
}

func (_ *Zakazchik) TableName() string {
	return "zakazchiki"
}

type StatusZakazchika struct {
	IdStatusa           uint   `json:"id_statusa" gorm:"primary_key;column:id_statusa"`
	NaimenovanieStatusa string `json:"naimenovanie_statusa" gorm:"column:naimenovanie_statusa"`
}

func (_ *StatusZakazchika) TableName() string {
	return "statusi_zakazchikov"
}
