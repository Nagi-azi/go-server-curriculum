package usecase

import (
	"go-server-curriculum/domain"
	"go-server-curriculum/repository"
)

type ProductUsecase struct {
	productRepo *repository.ProductRepository
}

// NewProductUsecase は ProductUsecase を初期化
func NewProductUsecase(productRepo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo: productRepo}
}

// GetAllProducts はすべての商品を取得
func (u *ProductUsecase) GetAllProducts() ([]domain.Product, error) {
	return u.productRepo.GetAllProducts()
}

// GetProductByID はIDで商品を取得
func (u *ProductUsecase) GetProductByID(id uint) (*domain.Product, error) {
	return u.productRepo.GetProductByID(id)
}

func (u *ProductUsecase) CreateProduct(product *domain.Product) error {
	return u.productRepo.CreateProduct(product)
}

func (u *ProductUsecase) UpdateProduct(id uint,product *domain.Product) error {
	new_product,result := u.productRepo.GetProductByID(id)
	new_product.Name = product.Name
	new_product.Price = product.Price

	if result != nil {	//nilは正常を意味する
		return result
	}

	result = u.productRepo.UpdateProduct(new_product)

	if result != nil {	//nilは正常を意味する
		return result
	}
	return nil
}

func (u *ProductUsecase) DeleteProduct(id uint) error {
	return u.productRepo.DeleteProduct(id)
}