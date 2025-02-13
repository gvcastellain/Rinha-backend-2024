package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gvcastellain/Rinha-backend-2024/db/sql/querys"
	"github.com/gvcastellain/Rinha-backend-2024/internal"
)

func ClientesHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	if len(paths) != 4 || paths[3] != "transacoes" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id := paths[2]

	_, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	transaction := internal.Transaction{}

	err = json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		return
	}

	if transaction.Type == "d" {
		exceeds, err := querys.ExceedsLimit(transaction, id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if exceeds {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
	}

	transactionResponse, err := querys.ExecTransaction(transaction, id)

	if err != nil {
		return
	}

	jsonData, err := json.Marshal(transactionResponse)

	if err != nil {
		return
	}

	w.Write([]byte(jsonData))
	w.WriteHeader(http.StatusOK)
}
