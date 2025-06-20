package handler

import (
	"net/http"
	"strconv"

	"go-server-curriculum/usecase"
	"go-server-curriculum/domain"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

// NewOrderHandler は OrderHandler を初期化
func NewOrderHandler(orderUsecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

// GetOrders はすべての注文を取得
func (h *OrderHandler) GetOrders(c echo.Context) error {
	orders, err := h.orderUsecase.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
	}
	return c.JSON(http.StatusOK, orders)
}

// GetOrder はIDで注文を取得
func (h *OrderHandler) GetOrder(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	order, err := h.orderUsecase.GetOrderByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}
	return c.JSON(http.StatusOK, order)
}

// CreateOrder は新しい注文を作成
func (h *OrderHandler) CreateOrder(c echo.Context) error {
		var order_data domain.Order
		if err := c.Bind(&order_data); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
	err := h.orderUsecase.CreateOrder(&order_data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	return c.JSON(http.StatusOK, "Create Order Success!!")
}

// UpdateOrder は注文を更新
func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))

	var order_data domain.Order
	if err := c.Bind(&order_data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := h.orderUsecase.UpdateOrder(uint(id), &order_data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusOK, "Update Order Success!!")
}

// DeleteOrder は注文を削除
func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))

	err := h.orderUsecase.DeleteOrder(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	return c.JSON(http.StatusOK, "Delete Order Success!!")
}