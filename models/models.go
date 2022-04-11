package models

import "gorm.io/gorm"

//Veiculo representa os dados do veículo em estoque
type Veiculo struct {
	gorm.Model            //já cria/atualiza os campos ID, CreatedAt, UpdatedAt, DeletedAt na DataBase
	Marca         string  `json:"marca"`
	Modelo        string  `json:"modelo"`
	Ano           string  `json:"ano"`
	Chassi        string  `json:"chassi"`
	Placa         string  `json:"placa"`
	Cor           string  `json:"cor"`
	ValorDeCompra float64 `json:"valor-de-compra"`
}
