package usecase

import (
	"go-server-curriculum/domain"
	"go-server-curriculum/repository"
)

type OrderUsecase struct {
	orderRepo *repository.OrderRepository
}

func NewOrderUsecase(orderRepo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: orderRepo}
}

func (u *OrderUsecase) GetAllOrders() ([]domain.Order, error) {
	return u.orderRepo.GetAllOrders()
}

func (u *OrderUsecase) GetOrderByID(id uint) (*domain.Order, error) {
	return u.orderRepo.GetOrderByID(id)
}

func (u *OrderUsecase) CreateOrder(order *domain.Order) error {
	return u.orderRepo.CreateOrder(order)
}

func (u *OrderUsecase) UpdateOrder(id uint, order *domain.Order) error {
	newOrder, err := u.orderRepo.GetOrderByID(id)
	if err != nil {
		return err
	}

	newOrder.ProductID = order.ProductID
	newOrder.Quantity = order.Quantity

	err = u.orderRepo.UpdateOrder(newOrder)
	if err != nil {
		return err
	}
	return nil
}

func (u *OrderUsecase) DeleteOrder(id uint) error {
	return u.orderRepo.DeleteOrder(id)
}