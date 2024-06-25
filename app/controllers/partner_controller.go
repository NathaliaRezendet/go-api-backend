package controllers

import (
    "net/http"
    "strconv"
    "go-api/app/models"
    "go-api/config"
    "github.com/gin-gonic/gin"
    "go-api/utils"
)

func GetPartners(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")       
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var partners []models.Partner

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&partners).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar parceiros"})
        return
    }

    c.JSON(http.StatusOK, partners)
}

func GetPartner(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var partner models.Partner
    if err := config.DB.First(&partner, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Parceiro não encontrado"})
        return
    }

    c.JSON(http.StatusOK, partner)
}

func CreatePartner(c *gin.Context) {
    var newPartner models.Partner
    if err := c.ShouldBindJSON(&newPartner); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&newPartner).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar parceiro"})
        return
    }

    c.JSON(http.StatusCreated, newPartner)
}

func UpdatePartner(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var updatedPartner models.Partner
    if err := c.ShouldBindJSON(&updatedPartner); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var partner models.Partner
    if err := config.DB.First(&partner, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Parceiro não encontrado"})
        return
    }

    updatedPartner.ID = partner.ID
    if err := config.DB.Save(&updatedPartner).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar parceiro"})
        return
    }

    c.JSON(http.StatusOK, updatedPartner)
}

func DeletePartner(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if err := config.DB.Delete(&models.Partner{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir parceiro"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Parceiro excluído com sucesso"})
}
