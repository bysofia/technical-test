package repositories

import (
	"nutech/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product models.Product) (models.Product,error)
	UpdateProduct(product models.Product, ID int) (models.Product, error)
	DeleteProduct(product models.Product, ID int) (models.Product, error)
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProduct(product models.Product) (models.Product,error){
	err := r.db.Create(&product).Error

	return product,err
}

func (r *repository) UpdateProduct(product models.Product, ID int) (models.Product,error){
	err := r.db.Model(&product).Where("id=?", ID).Updates(&product).Error

	return product,err
}

func (r *repository) DeleteProduct(product models.Product, ID int) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error

	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	
	err := r.db.First(&product, ID).Error

	return product, err
}