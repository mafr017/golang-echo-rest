package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/mafr017/golang-rest-echo/helpers"
	"github.com/mafr017/golang-rest-echo/models"
)

type JWTCustomClaims struct {
	Username string `json:"username"`
	Level string `json:"level"`
	jwt.StandardClaims
}

func GenerateHashPasswordControl(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLoginControl(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	/// set claims
	claims := JWTCustomClaims{
		username,
		"application",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	
	/// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/// Generate encoded token and send it as respone.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK,  echo.Map{
		"token": t,
	})

	/*
	/// create token
	token := jwt.New(jwt.SigningMethodHS256)

	/// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	/// Generate encoded token and send it as respone.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
	*/

	// return c.String(http.StatusOK, "Berhasil login..")
}