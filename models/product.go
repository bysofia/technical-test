package models

type Product struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Hargajual int    `json:"harga_jual" form:"harga_jual" gorm:"type: int"`
	Hargabeli int    `json:"harga_beli" form:"harga_beli" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock int    `json:"stock" form:"stock"`
}

type ProductResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Hargajual int    `json:"harga_jual" form:"harga_jual" gorm:"type: int"`
	Hargabeli int    `json:"harga_beli" form:"harga_beli" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock int    `json:"stock" form:"stock"`
}

func (ProductResponse) TableName() string {
	return "products"
}