package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mafr017/rest_echo/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil menggunakan echo")
	})
	e.GET("/pegawai", controllers.FetchAllPegawaiControl)
	e.POST("/pegawai", controllers.StorePegawaiControl)

	return e
}