package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-api/app/models"
    "go-api/config"
    "go-api/utils"
)

func GetClient(c *gin.Context) {
    var client models.Client
    id := c.Param("id")

    if err := config.DB.First(&client, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
        return
    }

    c.JSON(http.StatusOK, client)
}

func GetClients(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")      
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var clients []models.Client

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&clients).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
        return
    }

    c.JSON(http.StatusOK, clients)
}

func CreateClient(c *gin.Context) {
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&client).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": client.ID})
}

func UpdateClient(c *gin.Context) {
    var client models.Client
    id := c.Param("id")

    if err := config.DB.First(&client, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
        return
    }

    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&client).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cliente atualizado com sucesso"})
}

func DeleteClient(c *gin.Context) {
    var client models.Client
    id := c.Param("id")

    if err := config.DB.First(&client, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
        return
    }

    if err := config.DB.Delete(&client).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir cliente"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cliente excluído com sucesso"})
}
