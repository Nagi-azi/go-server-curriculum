package handler

import (
	"net/http"
	"strconv"

	"go-server-curriculum/usecase"
	"go-server-curriculum/domain"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerUsecase *usecase.CustomerUsecase
}

// NewCustomerHandler は CustomerHandler を初期化
func NewCustomerHandler(customerUsecase *usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{customerUsecase: customerUsecase}
}

// GetCustomers はすべての顧客を取得
func (h *CustomerHandler) GetCustomers(c echo.Context) error {
	customers, err := h.customerUsecase.GetAllCutomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch customers"})
	}
	return c.JSON(http.StatusOK, customers)
}

// GetCustomer はIDで顧客を取得
func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	customer, err := h.customerUsecase.GetCustomerByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Customer not found"})
	}
	return c.JSON(http.StatusOK, customer)
}

// CreateCustomer は新しい顧客を作成
func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err := h.customerUsecase.CreateCustomer(&customer)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	return c.JSON(http.StatusOK, "Create Customer Success!!")	
}

// UpdateCustomer は顧客を更新
func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	err = h.customerUsecase.UpdateCustomer(uint(id), &customer)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusOK, "Update Customer Success!!")
}

// DeleteCustomer は顧客を削除
func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	err = h.customerUsecase.DeleteCustomer(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusOK, "Delete Customer Success!!")
}

func (h *CustomerHandler) GetTotalPrice(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	customer, err := h.customerUsecase.GetTotalPrice(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Customer Total Price Not found"})
	}
	return c.JSON(http.StatusOK, customer)
}
