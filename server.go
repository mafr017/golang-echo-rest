package main

import (
	"github.com/mafr017/rest_echo/routes"
)

func main() {
	e := routes.Init()
	
	e.Logger.Fatal(e.Start(":1234"))
}