package models

type Application struct {
	IdZayavki           uint8          `json:"id_zayavki" gorm:"primaryKey;column:id_zayavki"`
	IdZakazchika        uint8          `json:"id_zakazchika" gorm:"column:id_zakazchika"`
	IdVidaPrazdnika     uint8          `json:"id_vida_prazdnika" gorm:"column:id_vida_prazdnika"`
	IdSotrudnika        uint8          `json:"id_sotrudnika" gorm:"column:id_sotrudnika"`
	IdStatusaZayavki    uint8          `json:"id_statusa_zayavki" gorm:"column:id_statusa_zayavki"`
	DataProvedeniya     string         `json:"data_provedeniya" gorm:"column:data_provedeniya"`
	KolichestvoChelovek string         `json:"kolichestvo_chelovek" gorm:"column:kolichestvo_chelovek"`
	NachaloProvedeniya  string         `json:"nachalo_provedeniya" gorm:"column:nachalo_provedeniya"`
	KonecProvedeniya    string         `json:"konec_provedeniya" gorm:"column:konec_provedeniya"`
	Prodoljitelnost     uint8          `json:"prodoljitelnost" gorm:"column:prodoljitelnost"`
	Zakazchik           Zakazchik      `gorm:"foreignKey:IdZakazchika;references:IdZakazchika"`
	VidiPrazdnikov      VidiPrazdnikov `gorm:"foreignKey:IdVidaPrazdnika;references:IdVida"`
	Sotrudnik           Sotrudnik      `json:"sotrudnik" gorm:"foreignKey:IdSotrudnika;references:IdSotrudnika"`
	StatusZayavki       StatusZayavki  `json:"status_zayavki" gorm:"foreignKey:IdStatusaZayavki;references:IdStatusa"`
}

func (_ *Application) TableName() string {
	return "zayavki"
}

type StatusZayavki struct {
	IdStatusa           uint   `json:"id_statusa_zayavki" gorm:"primary_key;column:id_statusa_zayavki"`
	NaimenovanieStatusa string `json:"naimenovanie_statusa_zayavki" gorm:"column:naimenovanie_statusa"`
}

func (_ *StatusZayavki) TableName() string {
	return "statusi_zayavok"
}

type SendApplication struct {
	IdVidaPrazdnika string   `json:"holidayId"`
	IdPolzovatelya  uint8    `json:"idPolzovatelya"`
	Kolichestvo     string   `json:"kolichestvo"`
	Date            string   `json:"date"`
	Nachalo         string   `json:"nachalo"`
	Konec           string   `json:"konec"`
	DopUslugi       []string `json:"additionalServices"`
}

type CustomerApplication struct {
	Prazdniki    []Prazdnik    `json:"prazdniki"`
	Applications []Application `json:"applications"`
}
