package controller

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Error struct {
	Message string `json:"message"`
}

var CustomerList = make(map[int]Customer)

var LastID int = 1

func GetCustomer(c echo.Context) (err error) {
	var result []Customer

	dataDummyCustomer := Customer{
		ID:    1,
		Name:  "Reni",
		Email: "reni@gmail.com",
	}

	result = append(result, dataDummyCustomer)

	for key, _ := range CustomerList {
		tempCustomer := CustomerList[key]

		result = append(result, tempCustomer)
	}

	return c.JSON(http.StatusOK, result)
}

func CreateCustomer(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)
	log.Println("Masuk ke create user")
	
	req := new(Customer)
	if err := c.Bind(req); err != nil {
		log.Println("Gagal bind")
		return c.JSON(http.StatusBadRequest, nil)
	}

	log.Println("Request data masuk", req)
	user := Customer{
		ID:    LastID,
		Name:  req.Name,
		Email: req.Email,
	}

	CustomerList[LastID] = user
	LastID++
	log.Println("Berhasil bind")

	return c.JSON(http.StatusCreated, user)
}

func GetCustomerByID(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)

	log.Println("Masuk ke get user by id")
	id, _ := strconv.Atoi(c.Param("id"))
	var result Customer
	
	for _, val := range CustomerList {
		if val.ID == id {
			result = val
		} else {
			log.Println("Gagal get user by id")
			return c.JSON(http.StatusNotFound, Error{
				Message: "Customer with id " + c.Param("id") + " not found",
			})
		}
	}
	log.Println("Berhasil get user by id")
	return c.JSON(http.StatusOK, result)
}


func UpdateCustomer(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)

	log.Println("Masuk ke update user")
	id, _ := strconv.Atoi(c.Param("id"))
	var result Customer
	
	for key, val := range CustomerList {
		if val.ID == id {
			result = val
			CustomerList[key] = Customer{
				ID:    val.ID,
				Name:  c.FormValue("name"),
				Email: c.FormValue("email"),
			}
		} else {
			log.Println("Gagal update user")
			return c.JSON(http.StatusBadRequest, Error{
				Message: "Customer with id " + c.Param("id") + " failed to update",
			})
		}
	}

	log.Println("Berhasil update user")
	return c.JSON(http.StatusOK, result)
}

func DeleteCustomer(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)

	id, _ := strconv.Atoi(c.Param("id"))
	var result Customer
	
	log.Println("Masuk ke delete user")
	for key, val := range CustomerList {
		if val.ID == id {
			result = val
			delete(CustomerList, key)
		} else {
			log.Println("Gagal delete user")
			return c.JSON(http.StatusNotFound, Error{
				Message: "Customer with id " + c.Param("id") + " not found",
			})
		}
	}
	log.Println("Berhasil delete user")
	return c.JSON(http.StatusOK, result)
}

