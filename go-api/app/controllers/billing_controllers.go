package controllers

import (
    "net/http"

    "go-api/app/models"
    "github.com/gin-gonic/gin"
    "go-api/config"
    "go-api/utils"
)

func GetBillings(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")      
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var billings []models.Billing

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&billings).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar detalhes de faturamento"})
        return
    }

    c.JSON(http.StatusOK, billings)
}

func GetBilling(c *gin.Context) {
    var billing models.Billing
    id := c.Param("id")

    if err := config.DB.First(&billing, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Detalhes de faturamento não encontrados"})
        return
    }

    c.JSON(http.StatusOK, billing)
}

func CreateBilling(c *gin.Context) {
    var billing models.Billing
    if err := c.ShouldBindJSON(&billing); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&billing).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar registro de faturamento"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Registro de faturamento criado com sucesso"})
}

func UpdateBilling(c *gin.Context) {
    var billing models.Billing
    id := c.Param("id")

    if err := config.DB.First(&billing, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Detalhes de faturamento não encontrados"})
        return
    }

    if err := c.ShouldBindJSON(&billing); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&billing).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar registro de faturamento"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registro de faturamento atualizado com sucesso"})
}

func DeleteBilling(c *gin.Context) {
    var billing models.Billing
    id := c.Param("id")

    if err := config.DB.First(&billing, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Detalhes de faturamento não encontrados"})
        return
    }

    if err := config.DB.Delete(&billing).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir registro de faturamento"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registro de faturamento excluído com sucesso"})
}
