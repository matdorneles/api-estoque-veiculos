package models

import "gorm.io/gorm"

//Veiculo representa os dados do veículo em estoque
type Veiculo struct {
	gorm.Model        //já cria/atualiza os campos ID, CreatedAt, UpdatedAt, DeletedAt na DataBase
	Marca      string `json:"marca"`
	Modelo     string `json:"modelo"`
	Ano        int    `json:"ano"`
	Cor        string `json:"cor"`
}
