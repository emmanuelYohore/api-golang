package main

import (
	"github.com/emmanuelYohore/api-golang/database"
	"github.com/emmanuelYohore/api-golang/routes"
)

func main() {
	database.Connect()
	routes.Routes()

}
