package models

type Ploshadka struct {
	IdPloshadki         uint8   `json:"id_ploshadki" gorm:"primaryKey;column:id_ploshadki"`
	Nazvanie            string  `json:"nazvanie" gorm:"column:nazvanie;size:50"`
	IdAddressaPloshadki uint8   `json:"id_addressa_ploshadki" gorm:"column:id_addressa_ploshadki;not null"`
	Vmestimost          string  `json:"vmestimost" gorm:"column:vmestimost;size:50"`
	StoimostVChas       float64 `json:"stoimost_v_chas" gorm:"column:stoimost_v_chas;size:20"`
	Address             Address `gorm:"foreignKey:IdAddressaPloshadki;references:IdAddressa"`
}

func (_ *Ploshadka) TableName() string {
	return "ploshadki"
}

type Address struct {
	IdAddressa uint8  `json:"id_addressa" gorm:"primaryKey;column:id_addressa"`
	Gorod      string `json:"gorod" gorm:"column:gorod;size:50"`
	Ulica      string `json:"ulica" gorm:"column:ulica;size:50"`
	Dom        string `json:"dom" gorm:"column:dom;size:50"`
	Korpus     string `json:"korpus" gorm:"column:korpus;size:50"`
}

func (_ *Address) TableName() string {
	return "addressa_ploshadok"
}
