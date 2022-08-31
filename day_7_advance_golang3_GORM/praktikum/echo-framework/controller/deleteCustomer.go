package controller

import (
	"echo-framework/model"
	"net/http"
	"log"
	"os"
	"github.com/labstack/echo/v4"
)

func (h Handler) DeleteCustomer(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)

	log.Println("Masuk ke delete user")
	var customer model.Customer
	if err := c.Bind(&customer); err != nil {
		return err
	}
	h.db.Delete(&customer)
	log.Println("Berhasil delete user")
	
	return c.JSON(http.StatusOK, customer)
}