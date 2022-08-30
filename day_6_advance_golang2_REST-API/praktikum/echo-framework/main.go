package main

import (
	"echo-framework/controller"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	app_name := os.Getenv("APP_NAME")
	if app_name 
	// initialize echo framework
	e := echo.New()

	// root
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message":       "Success to connect service",
		}

		return c.JSON(http.StatusOK, result)
	})

	cust := e.Group("customer")
	cust.GET("", controller.GetCustomer)
	cust.GET("/:id", controller.GetCustomerByID)
	cust.POST("", controller.CreateCustomer)
	cust.PUT("/:id", controller.UpdateCustomer)
	cust.DELETE("/:id", controller.DeleteCustomer)

	// start service echo
	e.Logger.Fatal(e.Start(":5002"))
}
