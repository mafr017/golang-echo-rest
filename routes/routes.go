package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mafr017/golang-rest-echo/controllers"
	"github.com/mafr017/golang-rest-echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil menggunakan echo")
	})

	// Simple CRUD
	e.GET("/pegawai", controllers.FetchAllPegawaiControl, middleware.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawaiControl)
	e.PUT("/pegawai", controllers.UpdatePegawaiControl)
	e.DELETE("/pegawai", controllers.DeletePegawaiControl)

	// User Login Routes
	e.GET("/generate-hash/:password", controllers.GenerateHashPasswordControl)
	e.POST("/login", controllers.CheckLoginControl)

	return e
}