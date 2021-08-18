package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	Nama string `validate:"required"`
	Email string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur int `validate:"required,gte=17,lte=45"`
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama: "adit",
		Email: "adit@email.com",
		Alamat: "jl.cihanjuang",
		Umur: 23,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Validation success..",
	})
}

func TestVariableValidation(c echo.Context) error {
	v := validator.New()

	email :=  "aditt@email.com"

	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "email tidak valid",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Validation email success..",
	})
}