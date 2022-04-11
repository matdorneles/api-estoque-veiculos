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

//exibe veículo pelo ID
func ExibeVeiculoId(c *gin.Context) {
	var veiculo models.Veiculo
	id := c.Params.ByName("id")
	database.DB.First(&veiculo, id)

	if veiculo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Não Encontrado": "Veículo não encontrado"})
		return
	}
	c.JSON(http.StatusOK, veiculo)
}

//exibe veículo pela Marca
func ExibeVeiculoMarca(c *gin.Context) {
	var veiculo []models.Veiculo
	marca := c.Param("marca")
	database.DB.Where(&models.Veiculo{Marca: marca}).Find(&veiculo)

	c.JSON(http.StatusOK, veiculo)
}

func ExibeVeiculoModelo(c *gin.Context) {
	var veiculo []models.Veiculo
	modelo := c.Param("modelo")
	database.DB.Where(&models.Veiculo{Modelo: modelo}).Find(&veiculo)

	c.JSON(http.StatusOK, veiculo)
}

func ExibeVeiculoCor(c *gin.Context) {
	var veiculo []models.Veiculo
	cor := c.Param("cor")
	database.DB.Where(&models.Veiculo{Cor: cor}).Find(&veiculo)

	c.JSON(http.StatusOK, veiculo)
}

//cria um novo veículo
func CriaVeiculo(c *gin.Context) {
	var veiculo models.Veiculo
	if err := c.ShouldBindJSON(&veiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	} else if len(veiculo.Chassi) != 17 { //se a quantidade de caracteres em chassi for diferente que 17, não é possível criar o veículo
		c.JSON(http.StatusBadGateway, gin.H{
			"Erro": "O Chassi deve conter 17 caracteres"})
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
