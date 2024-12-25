package models

type Vedushiy struct {
	IdVedushego uint8   `json:"id_vedushego" gorm:"primaryKey;column:id_vedushego"`
	Familiya    string  `json:"familiya" gorm:"column:familiya;size:50;not null"`
	Imya        string  `json:"imya" gorm:"column:imya;size:50;not null"`
	Otchestvo   string  `json:"otchestvo" gorm:"column:otchestvo;size:50"`
	Telephone   string  `json:"telephone" gorm:"column:telephone;size:12;not null"`
	Email       string  `json:"email" gorm:"column:email;not null"`
	Photo       string  `json:"photo" gorm:"column:photo;not null"`
	StavkaVChas float64 `json:"stavka_v_chas" gorm:"column:stavka_v_chas"`
}

func (_ *Vedushiy) TableName() string {
	return "vedushie"
}
