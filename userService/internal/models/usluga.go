package models

type UslugiVZayavkah struct {
	IdUslugiVZayavke uint        `json:"id_uslugi_v_zayavke" gorm:"primaryKey;column:id_uslugi_v_zayavke"`
	IdUslugi         uint        `json:"id_uslugi" gorm:"column:id_uslugi;not null"`
	IdZayavki        uint        `json:"id_zayavki" gorm:"column:id_zayavki;not null"`
	Complete         bool        `json:"complete" gorm:"column:complete;not null"`
	Usluga           Usluga      `gorm:"foreignKey:IdUslugi;references:IdUslugi"`
	Zayavka          Application `gorm:"foreignKey:IdZayavki;references:IdZayavki"`
}

func (_ *UslugiVZayavkah) TableName() string {
	return "uslugi_v_zayavkah"
}

type Usluga struct {
	IdUslugi     uint    `json:"id_uslugi" gorm:"primaryKey;column:id_uslugi"`
	Naimenovanie string  `json:"naimenovanie" gorm:"column:naimenovanie;size:50;not null"`
	Stoimost     float64 `json:"stoimost" gorm:"column:stoimost;not null"`
	Opisanie     string  `json:"opisanie" gorm:"column:opisanie;size:50;not null"`
}

func (_ *Usluga) TableName() string {
	return "uslugi"
}
