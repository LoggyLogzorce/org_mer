package models

type Event struct {
	IdZayavki           uint8  `json:"id_zayavki" gorm:"primary_key"`
	Familiya            string `json:"familiya_zakazchika"`
	Imya                string `json:"imya_zakazchika"`
	Otchestvo           string `json:"otchestvo_zakazchika"`
	NaimenovanieStatusa string `json:"stasus_zakazchika"`
	Telephone           string `json:"telephone_zakazchika"`
	Email               string `json:"email_zakazchika"`
	NaimenovanieVida    string `json:"vid_prazdnika"`
	DataProvedeniya     string `json:"data_provedeniya"`
	KolichestvoChelovek string `json:"kolichestvo_chelovek"`
	NachaloProvedeniya  string `json:"nachalo_provedeniya"`
	KonecProvedeniya    string `json:"konec_provedeniya"`
	Prodoljitelnost     uint8  `json:"prodoljitelnost"`
}
