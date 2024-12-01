package models

type Event struct {
	IdZayavki           uint8  `json:"id_zayavki" gorm:"primary_key"`
	Familiya            string `json:"familiya_zakazckika"`
	Imya                string `json:"imya_zakazchika"`
	Otchestvo           string `json:"otchestvo_zakazchika"`
	NaimenovanieStatusa string `json:"stasus_zakazchika"`
	Telephone           string `json:"telephone_zakazchika"`
	Email               string `json:"email_zakazchika"`
	NaimenovanieVida    string `json:"vid_prazdnika"`
	DataProvedeniya     string `json:"data_provedeniya"`
}

func (_ *Event) TableName() string {
	return "zayavki"
}
