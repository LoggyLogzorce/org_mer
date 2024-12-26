package models

type VidiPrazdnikov struct {
	IdVida           uint8   `json:"id_vida" gorm:"primaryKey;column:id_vida"`
	NaimenovanieVida string  `json:"naimenovanie_vida" gorm:"column:naimenovanie_vida"`
	Summa            float64 `json:"summa" gorm:"column:summa"`
	Opisanie         string  `json:"opisanie" gorm:"column:opisanie"`
	Photo            string  `json:"photo" gorm:"column:photo"`
}

func (_ *VidiPrazdnikov) TableName() string {
	return "vidi_prazdnikov"
}
