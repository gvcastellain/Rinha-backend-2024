package internal

type Transaction struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo" validate:"oneof=c d"`
	Description string `json:"descricao" validate:"min=1,max=10"`
}

type TransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}
