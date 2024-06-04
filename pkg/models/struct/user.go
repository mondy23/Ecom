package models

type User struct {
	Userid   uint     `json:"userid" gorm:"primarykey"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Token    string   `json:"token"`
	Customer Customer `json:"customer"`
}
