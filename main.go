package main

import (
	"github.com/matdorneles/estoque-veiculos/database"
	"github.com/matdorneles/estoque-veiculos/routes"
)

func main() {
	//Iniciando conexão com database PostgreSQL utilizando GORM, lembre-se de inicializar o docker antes
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
