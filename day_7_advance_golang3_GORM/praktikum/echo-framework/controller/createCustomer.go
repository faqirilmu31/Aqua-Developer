package controller

import (
	"echo-framework/model"
	"net/http"
	"log"
	"os"
	"github.com/labstack/echo/v4"
)

func (h Handler) CreateCustomer(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)

	log.Println("Masuk ke create user")
	var customer model.Customer
	req := new(model.Customer)
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := model.Customer{
		ID:    model.LastID,
		Name:  req.Name,
		Email: req.Email,
	}

	model.CustomerList[model.LastID] = user
	model.LastID++

	if result := h.db.Create(&customer); result.Error != nil {
		return c.JSON(http.StatusBadRequest, model.Error{
			Message: "Gagal create user",
		})
	}
	log.Println("Berhasil create user")

	return c.JSON(http.StatusCreated, customer)
}