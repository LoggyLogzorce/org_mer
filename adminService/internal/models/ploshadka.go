package models

type Ploshadka struct {
	IdPloshadki         uint8   `json:"id_ploshadki" gorm:"primaryKey;column:id_ploshadki"`
	Nazvanie            string  `json:"nazvanie" gorm:"column:nazvanie;size:50"`
	IdAddressaPloshadki uint8   `json:"id_addressa_ploshadki" gorm:"column:id_addressa_ploshadki;not null"`
	Vmestimost          string  `json:"vmestimost" gorm:"column:vmestimost;size:50"`
	StoimostVChas       float64 `json:"stoimost_v_chas" gorm:"column:stoimost_v_chas;size:20"`
	//IdTematiki       uint8   `json:"id_tematiki" gorm:"column:id_tematiki;not null"`
	//AddressaPloshadki AddressaPloshadok `gorm:"foreignKey:IdAddressaPloshadki;references:IdAddressa"`
	//TematikaPloshadki TematikaPloshadok `gorm:"foreignKey:IdTematiki;references:IdTematiki"`
}

func (_ *Ploshadka) TableName() string {
	return "ploshadki"
}
