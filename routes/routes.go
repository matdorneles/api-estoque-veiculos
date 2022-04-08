package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matdorneles/estoque-veiculos/controllers"
)

func HandleRequests() {
	r := gin.Default() //facilita a descrição das rotas abaixo
	r.GET("/estoque", controllers.ExibeTodosVeiculos)
	r.POST("/estoque/criar", controllers.CriaVeiculo)
	r.PATCH("/estoque/alterar/:id", controllers.EditaDadosVeiculo)
	r.DELETE("/estoque/deletar/:id", controllers.DeletaVeiculo)
	r.Run()
}
