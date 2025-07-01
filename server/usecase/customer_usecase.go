package usecase

import (
	"go-server-curriculum/domain"
	"go-server-curriculum/repository"
)

type CustomerUsecase struct {
	customerRepo *repository.CustomerRepository
	productRepo  *repository.ProductRepository
}

// NewCustomerUsecase は CustomerUsecase を初期化
func NewCustomerUsecase(customerRepo *repository.CustomerRepository,productRepo  *repository.ProductRepository) *CustomerUsecase {
	return &CustomerUsecase{
		customerRepo: customerRepo,
		productRepo:  productRepo,
	}
}

// GetAllCutomers はすべての顧客を取得
func (u *CustomerUsecase) GetAllCutomers() ([]domain.Customer, error) {
	return u.customerRepo.GetAllCustomers()
}

// GetCustomerByID はIDで顧客を取得
func (u *CustomerUsecase) GetCustomerByID(id uint) (*domain.Customer, error) {
	customer, err := u.customerRepo.GetCustomerByID(id)
	return customer, err
}

func (u *CustomerUsecase) CreateCustomer(customer *domain.Customer) error {
	return u.customerRepo.CreateCustomer(customer)
}

func (u *CustomerUsecase) UpdateCustomer(id uint, customer *domain.Customer) error {
	newCustomer, err := u.customerRepo.GetCustomerByID(id)
	if err != nil {
		return err
	}

	newCustomer.Name = customer.Name
	newCustomer.Seat = customer.Seat

	err = u.customerRepo.UpdateCustomer(newCustomer)
	if err != nil {
		return err
	} 

	return  nil
}

func (u *CustomerUsecase) DeleteCustomer(id uint) error {
	return u.customerRepo.DeleteCustomer(id)
}

func (u *CustomerUsecase) GetTotalPrice(id uint) (map[string]interface{}, error) {
	customer, err := u.customerRepo.GetCustomerWithOrders(id)
	if err != nil {
		return nil, err
	}

	total := 0
	for _, order := range customer.Orders {
		product, err := u.productRepo.GetProductByID(order.ProductID)
		if err != nil {
			return nil, err
		}

		total += order.Quantity * product.Price
	}

	return map[string]interface{}{
		"customer" : customer,
		"total" : total,
 	}, nil

}