package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matdorneles/estoque-veiculos/controllers"
)

func HandleRequests() {
	r := gin.Default() //iniciando variavel de rotas Gin
	r.GET("/estoque", controllers.ExibeTodosVeiculos)
	r.GET("/estoque/id/:id", controllers.ExibeVeiculoId)
	r.GET("/estoque/marca/:marca", controllers.ExibeVeiculoMarca)
	r.GET("/estoque/modelo/:modelo", controllers.ExibeVeiculoModelo)
	r.GET("/estoque/cor/:cor", controllers.ExibeVeiculoCor)
	r.POST("/estoque/criar", controllers.CriaVeiculo)
	r.PATCH("/estoque/alterar/:id", controllers.EditaDadosVeiculo)
	r.DELETE("/estoque/deletar/:id", controllers.DeletaVeiculo)
	r.Run()
}
