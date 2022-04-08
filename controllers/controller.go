package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matdorneles/estoque-veiculos/database"
	"github.com/matdorneles/estoque-veiculos/models"
)

//exibe todos os veículos cadastrados
func ExibeTodosVeiculos(c *gin.Context) {
	var veiculos []models.Veiculo
	database.DB.Find(&veiculos)
	c.JSON(200, veiculos)
}

//cria um novo veículo
func CriaVeiculo(c *gin.Context) {
	var veiculo models.Veiculo
	if err := c.ShouldBindJSON(&veiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&veiculo)
	c.JSON(http.StatusOK, veiculo)
}

//altera dados de um veículo pelo ID
func EditaDadosVeiculo(c *gin.Context) {
	var veiculo models.Veiculo
	id := c.Params.ByName("id")
	database.DB.First(&veiculo, id)

	if err := c.ShouldBindJSON(&veiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&veiculo).UpdateColumns(veiculo)
	c.JSON(http.StatusOK, veiculo)
}

//deleta veículo pelo ID
func DeletaVeiculo(c *gin.Context) {
	var veiculo models.Veiculo
	id := c.Params.ByName("id")
	database.DB.Delete(&veiculo, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Veículo deletado com sucesso"})
}