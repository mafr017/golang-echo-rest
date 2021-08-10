package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mafr017/rest_echo/models"
)

func FetchAllPegawaiControl(c echo.Context) error {
	ress, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, ress)
}