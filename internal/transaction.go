package internal

type Transaction struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo" validate:"oneof=c d"`
	Description string `json:"descricao" validate:"min=1,max=10"`
}
