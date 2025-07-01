package main

import (
	"go-server-curriculum/handler"
	"log"

	"go-server-curriculum/infrastructure"
	"go-server-curriculum/repository"
	"go-server-curriculum/usecase"
	
	"github.com/labstack/echo/v4"
)

func main() {
	// DB 初期化
	db, err := infrastructure.NewMySQLDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// // リポジトリ初期化
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	customerRepo := repository.NewCustomerRepository(db)

	// // ユースケース初期化
	productUsecase := usecase.NewProductUsecase(productRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo, productRepo)

	// // ハンドラー初期化
	healthHandler := handler.NewHealthHandler()
	productHandler := handler.NewProductHandler(productUsecase)
	orderHandler := handler.NewOrderHandler(orderUsecase)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	// Echoルーター設定
	e := echo.New()
	e.GET("/", healthHandler.HealthCheck)

	// product
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProduct)
	e.POST("/products", productHandler.CreateProduct) //create POST
	e.PUT("/products/:id", productHandler.UpdateProduct) //update PUT
	e.DELETE("/products/:id", productHandler.DeleteProduct) //delete DELETE

	// order
	e.GET("/orders", orderHandler.GetOrders)
	e.GET("/orders/:id", orderHandler.GetOrder)
	e.POST("/orders", orderHandler.CreateOrder)// create POST
	e.PUT("/orders/:id", orderHandler.UpdateOrder) // update PUT
	e.DELETE("/orders/:id", orderHandler.DeleteOrder) // delete DELETE

	// customer
	e.GET("/customers", customerHandler.GetCustomers)
	e.GET("/customers/:id", customerHandler.GetCustomer)
	e.POST("/customers", customerHandler.CreateCustomer) // create POST
	e.PUT("/customers/:id", customerHandler.UpdateCustomer) // update PUT
	e.DELETE("/customers/:id", customerHandler.DeleteCustomer) // delete DELETE
	e.GET("/customers/:id/total", customerHandler.GetTotalPrice) // total price GET

	// サーバー起動
	log.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
