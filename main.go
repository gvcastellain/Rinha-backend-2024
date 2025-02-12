package main

import (
	"github.com/gvcastellain/Rinha-backend-2024/db/connection"
	routes "github.com/gvcastellain/Rinha-backend-2024/routes"
)

func main() {
	connection.NewDbConn()
	routes.CreateRoutes()
}
