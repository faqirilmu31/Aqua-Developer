package controller

import "gorm.io/gorm"

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// func (h *Handler) GetCustomer(c echo.Context) error {
// 	var result []Customer
// 	var customer Customer
// 	var err error

// 	h.db.First(&customer)
// 	result = append(result, customer)
// 	return c.JSON(http.StatusOK, result)
// }