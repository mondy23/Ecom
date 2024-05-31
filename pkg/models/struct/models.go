package models

type Product struct {
	// Don't use camelcase on naming struct
	Productid   uint   `json:"productid" gorm:"primarykey"`
	Productname string `json:"productname"`
	Supplierid  uint   `json:"supplierid"`
	Categoryid  uint   `json:"categoryid"`
	Unit        string `json:"unit"`
	Price       int    `json:"price"`
}
