package models

import (
	"time"
)

// to exclude to response use -
type Product struct {
	// Don't use camelcase on naming struct
	Productid   uint   `json:"productid" gorm:"primarykey"`
	Productname string `json:"productname"`
	Supplierid  uint   `json:"supplierid"`
	Categoryid  uint   `json:"categoryid"`
	Unit        string `json:"unit"`
	Price       int    `json:"price"`
}

type Customer struct {
	Customerid   uint   `json:"customerid" gorm:"primarykey"`
	Customername string `json:"customername"`
	Contactname  string `json:"contactname"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Postalcode   string `json:"postalcode"`
	Country      string `json:"country"`
}

type Order struct {
	Orderid     uint `json:"orderid" gorm:"primarykey"`
	Customerid  uint `json:"customerid"`
	Employeeid  uint `json:"-"`
	Shipperid   uint `json:"shipperid"`
	Orderdate   time.Time
	Orderdetail Order_detail `gorm:"foreignkey:orderid"`
}

type Order_detail struct {
	Orderdetailid uint `gorm:"primaryKey"`
	Orderid       uint
	Productid     uint
	Quantity      int
	Product       Product `gorm:"foreignkey:productid"`
}
