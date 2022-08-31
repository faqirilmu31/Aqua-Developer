package controller

import (
	"echo-framework/model"
	"net/http"
	"log"
	"os"
	"strconv"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetCustomerByID(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)
	log.Println("Masuk ke get user by ID")

	var result []model.Customer
	var customer model.Customer
	id, _ := strconv.Atoi(c.Param("id"))

	log.Println("Request data masuk dgn ID", id)
	h.db.First(&customer, id)
	result = append(result, customer)

	log.Println("Berhasil get user by ID")
	return c.JSON(http.StatusOK, result)
}