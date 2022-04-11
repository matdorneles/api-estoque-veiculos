package database

import (
	"log"

	"github.com/matdorneles/estoque-veiculos/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Declaração de variáveis para a base de dados
var (
	DB  *gorm.DB
	err error
)

//Conectando à base de dados PostgreSQL utilizando GORM
func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Veiculo{})
}
