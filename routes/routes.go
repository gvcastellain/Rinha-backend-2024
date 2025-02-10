package routes

import (
	"net/http"

	"github.com/gvcastellain/Rinha-backend-2024/handlers"
)

func CreateRoutes() {
	http.HandleFunc("POST /clientes/", handlers.ClientesHandler)
	http.ListenAndServe(":9999", nil)
}
