package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-api/app/models"
    "go-api/config"
    "go-api/utils"
)

func GetBenefitEntities(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")      
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var benefits []models.BenefitEntity

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&benefits).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar benefícios"})
        return
    }

    c.JSON(http.StatusOK, benefits)
}

func GetBenefitEntity(c *gin.Context) {
    var benefit models.BenefitEntity
    id := c.Param("id")

    if err := config.DB.First(&benefit, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Benefício não encontrado"})
        return
    }

    c.JSON(http.StatusOK, benefit)
}

func CreateBenefitEntity(c *gin.Context) {
    var benefit models.BenefitEntity
    if err := c.ShouldBindJSON(&benefit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&benefit).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar benefício"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": benefit.ID})
}

func UpdateBenefitEntity(c *gin.Context) {
    var benefit models.BenefitEntity
    id := c.Param("id")

    if err := config.DB.First(&benefit, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Benefício não encontrado"})
        return
    }

    if err := c.ShouldBindJSON(&benefit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&benefit).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar benefício"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Benefício atualizado com sucesso"})
}

func DeleteBenefitEntity(c *gin.Context) {
    var benefit models.BenefitEntity
    id := c.Param("id")

    if err := config.DB.First(&benefit, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Benefício não encontrado"})
        return
    }

    if err := config.DB.Delete(&benefit).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir benefício"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Benefício excluído com sucesso"})
}
