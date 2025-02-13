package querys

import (
	"github.com/gvcastellain/Rinha-backend-2024/db/connection"
	"github.com/gvcastellain/Rinha-backend-2024/internal"
)

func ExceedsLimit(transaction internal.Transaction, id string) (bool, error) {
	db := connection.GetDB()

	defer db.Close()

	query := "SELECT current_limit - current_balance from client where id = ?"

	var balance int
	err := db.QueryRow(query, id).Scan(&balance)

	if err != nil {
		return true, err
	}

	if transaction.Value > balance {
		return true, nil
	}

	return false, nil
}

func ExecTransaction(transaction internal.Transaction, id string) (response *internal.TransactionResponse, err error) {
	db := connection.GetDB()

	query := "SELECT client_limit, current_balance from client where id = $1"

	var limit, balance int
	_ = db.QueryRow(query, id).Scan(&limit, &balance)

	balance += transaction.Value

	query = "UPDATE client SET current_balance = $1 WHERE id = $2"

	_, err = db.Exec(query, balance, id)

	if err != nil {
		return nil, err
	}

	query = "INSERT INTO transaction (client_id, description, value, transaction_type) VALUES ($1,$2,$3,$4)"

	_, err = db.Exec(query, id, transaction.Description, transaction.Value, transaction.Type)

	if err != nil {
		return nil, err
	}

	return &internal.TransactionResponse{Limit: limit, Balance: balance}, nil
}
