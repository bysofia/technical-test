package productdto

type ProductRequest struct {
	Name         string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Pricebuying  int    `json:"price_buying" form:"price_buying" gorm:"type: int"`
	Priceselling int    `json:"price_selling" form:"price_selling" gorm:"type: int"`
	Image        string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock        int    `json:"stock" form:"stock" gorm:"type: int"`
}

type UpdateProductResponse struct {
	Name         string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Pricebuying  int    `json:"price_buying" form:"price_buying" gorm:"type: int"`
	Priceselling int    `json:"price_selling" form:"price_selling" gorm:"type: int"`
	Image        string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock        int    `json:"stock" form:"stock" gorm:"type: int"`
}