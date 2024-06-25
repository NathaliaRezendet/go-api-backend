package controllers

import (
    "net/http"
    "go-api/app/models"
    "github.com/gin-gonic/gin"
    "go-api/config"
    "go-api/utils"
)

func GetResourceUsages(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")        
    pageSizeStr := c.DefaultQuery("page_size", "10")

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var resourceUsages []models.ResourceUsage

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&resourceUsages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usos de recursos"})
        return
    }

    c.JSON(http.StatusOK, resourceUsages)
}

func GetResourceUsage(c *gin.Context) {
    id := c.Param("id")
    var resourceUsage models.ResourceUsage

    if err := config.DB.First(&resourceUsage, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Uso de recurso não encontrado"})
        return
    }

    c.JSON(http.StatusOK, resourceUsage)
}

func CreateResourceUsage(c *gin.Context) {
    var newResourceUsage models.ResourceUsage

    if err := c.ShouldBindJSON(&newResourceUsage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&newResourceUsage).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar uso de recurso"})
        return
    }

    c.JSON(http.StatusCreated, newResourceUsage)
}

func UpdateResourceUsage(c *gin.Context) {
    id := c.Param("id")
    var updatedResourceUsage models.ResourceUsage

    if err := c.ShouldBindJSON(&updatedResourceUsage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&models.ResourceUsage{}).Where("id = ?", id).Updates(&updatedResourceUsage).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar uso de recurso"})
        return
    }

    c.JSON(http.StatusOK, updatedResourceUsage)
}

func DeleteResourceUsage(c *gin.Context) {
    id := c.Param("id")
    var resourceUsage models.ResourceUsage

    if err := config.DB.First(&resourceUsage, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Uso de recurso não encontrado"})
        return
    }

    if err := config.DB.Delete(&resourceUsage).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir uso de recurso"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Uso de recurso excluído com sucesso"})
}
