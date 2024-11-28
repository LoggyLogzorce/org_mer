package user

type User struct {
	Uid      uint8  `json:"uid" gorm:"primaryKey"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Username string `json:"username"`
	Role     uint8  `json:"role"`
}
