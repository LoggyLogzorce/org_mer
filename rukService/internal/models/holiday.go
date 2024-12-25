package models

type Prazdnik struct {
	IdPrazdnika     uint8       `json:"id_prazdnika" gorm:"primaryKey;column:id_prazdnika"`
	IdZayavki       uint8       `json:"id_zayavki" gorm:"column:id_zayavki"`
	IdVedushevo     uint8       `json:"id_vedushevo" gorm:"column:id_vedushevo"`
	IdPloshadki     uint8       `json:"id_ploshadki" gorm:"column:id_ploshadki"`
	PolnayaStoimost float64     `json:"polnaya_stoimost" gorm:"column:polnaya_stoimost"`
	Zayavka         Application `gorm:"foreignKey:IdZayavki;references:IdZayavki"`
	Vedushiy        Vedushiy    `gorm:"foreignKey:IdVedushevo;references:IdVedushego"`
	Ploshadka       Ploshadka   `gorm:"foreignKey:IdPloshadki;references:IdPloshadki"`
}

func (_ *Prazdnik) TableName() string {
	return "prazdniki"
}

type Holiday struct {
	Prazdnik  Prazdnik
	Ploshadki []Ploshadka
	Statusi   []StatusZayavki
	Vedushie  []Vedushiy
	Uslugi    []UslugiVZayavkah
}

type HolidayData struct {
	IDZayavki           string   `json:"id_zayavki"`
	IDVidaPrazdnika     string   `json:"id_vida_prazdnika"`
	DataProvedeniya     string   `json:"data_provedeniya"`
	KolichestvoChelovek string   `json:"kolichestvo_chelovek"`
	NachaloProvedeniya  string   `json:"nachalo_provedeniya"`
	KonecProvedeniya    string   `json:"konec_provedeniya"`
	Prodoljitelnost     string   `json:"prodoljitelnost"`
	IDStatusaZayavki    string   `json:"id_statusa_zayavki"`
	IDPloshadki         string   `json:"id_ploshadki"`
	IDVedushego         string   `json:"id_vedushego"`
	DopUslugi           []string `json:"dop_uslugi"`
}
