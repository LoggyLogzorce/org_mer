package models

type User struct {
	IDPolzovatelya uint8            `gorm:"primaryKey;"`
	Login          string           `gorm:"not null"`
	Password       string           `gorm:"not null"`
	IDRoli         uint             `gorm:"not null"`
	Rola           RolaPolzovatelya `gorm:"foreignKey:IDRoli;references:IDRoli"`
}

func (_ *User) TableName() string {
	return "polzovateli"
}

type RolaPolzovatelya struct {
	IDRoli           uint   `gorm:"primaryKey;autoIncrement"`
	NaimenovanieRoli string `gorm:"not null"`
}

func (_ *RolaPolzovatelya) TableName() string {
	return "roli_polzovateley"
}
