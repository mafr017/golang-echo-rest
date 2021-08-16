package main

import (
	"github.com/mafr017/golang-rest-echo/db"
	"github.com/mafr017/golang-rest-echo/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}