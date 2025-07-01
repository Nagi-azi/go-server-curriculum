package	repository

import (
	"go-server-curriculum/domain"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository は CustomerRepository を初期化
func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// GetAllCustomers はすべての顧客を取得
func (r *CustomerRepository) GetAllCustomers() ([]domain.Customer, error) {
	var customers []domain.Customer
	result := r.db.Find(&customers)
	return customers, result.Error
}

// GetCustomerByID はIDから顧客を取得
func (r *CustomerRepository) GetCustomerByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	result := r.db.First(&customer, id)
	return &customer, result.Error
}

// CreateCustomer は新しい顧客を作成
func (r *CustomerRepository) CreateCustomer(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

// UpdateCustomer は既存の顧客を更新
func (r *CustomerRepository) UpdateCustomer(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

// DeleteCustomer は顧客を削除
func (r *CustomerRepository) DeleteCustomer(id uint) error {
	return r.db.Delete(&domain.Customer{}, id).Error
}

func (r *CustomerRepository) GetCustomerWithOrders(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	result := r.db.Preload("Orders").First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}