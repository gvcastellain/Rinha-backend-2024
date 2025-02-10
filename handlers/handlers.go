package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gvcastellain/Rinha-backend-2024/internal"
)

func ClientesHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	if len(paths) != 4 || paths[3] != "transacoes" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	urlId := paths[2]

	_, err := strconv.Atoi(urlId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return
	}

	transaction := internal.Transaction{}

	err = json.Unmarshal(body, &transaction)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("a"))
}
