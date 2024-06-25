package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-api/app/models"
    "go-api/config"
    "go-api/utils"
)

func GetTag(c *gin.Context) {
    var tag models.Tag
    id := c.Param("id")

    if err := config.DB.First(&tag, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tag não encontrada"})
        return
    }

    c.JSON(http.StatusOK, tag)
}

func GetTags(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")         
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var tags []models.Tag

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&tags).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tags"})
        return
    }

    c.JSON(http.StatusOK, tags)
}

func CreateTag(c *gin.Context) {
    var tag models.Tag
    if err := c.ShouldBindJSON(&tag); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&tag).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar tag"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": tag.ID})
}

func UpdateTag(c *gin.Context) {
    var tag models.Tag
    id := c.Param("id")

    if err := config.DB.First(&tag, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tag não encontrada"})
        return
    }

    if err := c.ShouldBindJSON(&tag); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&tag).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar tag"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Tag atualizada com sucesso"})
}

func DeleteTag(c *gin.Context) {
    var tag models.Tag
    id := c.Param("id")

    if err := config.DB.First(&tag, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tag não encontrada"})
        return
    }

    if err := config.DB.Delete(&tag).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar tag"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Tag deletada com sucesso"})
}
