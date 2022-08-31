package controller

import (
	"echo-framework/model"
	"net/http"
	"log"
	"os"
	"strconv"
	"github.com/labstack/echo/v4"
)

func (h Handler) UpdateCustomer(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)
	log.Println("Masuk ke update user")

	id, _ := strconv.Atoi(c.Param("id"))
	var customer model.Customer
	if err := c.Bind(&customer); err != nil {
		return err
	}

	h.db.Model(&customer).Where("id = ?", id).Update("name", customer.Name)
	h.db.Model(&customer).Where("id = ?", id).Update("email", customer.Email)
	log.Println("Berhasil update user")

	return c.JSON(http.StatusOK, customer)
}