package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-api/app/models"
    "go-api/config"
    "go-api/utils"
)

func GetServiceInfo(c *gin.Context) {
    id := c.Param("id")
    var serviceInfo models.ServiceInfo

    if err := config.DB.First(&serviceInfo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Informações do serviço não encontradas"})
        return
    }

    c.JSON(http.StatusOK, serviceInfo)
}

func GetServiceInfos(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")          
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var serviceInfos []models.ServiceInfo

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&serviceInfos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar informações do serviço"})
        return
    }

    c.JSON(http.StatusOK, serviceInfos)
}

func CreateServiceInfo(c *gin.Context) {
    var newServiceInfo models.ServiceInfo
    if err := c.ShouldBindJSON(&newServiceInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&newServiceInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar informações do serviço"})
        return
    }

    c.JSON(http.StatusCreated, newServiceInfo)
}

func UpdateServiceInfo(c *gin.Context) {
    id := c.Param("id")
    var updatedServiceInfo models.ServiceInfo
    if err := c.ShouldBindJSON(&updatedServiceInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var existingServiceInfo models.ServiceInfo
    if err := config.DB.First(&existingServiceInfo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Serviço não encontrado"})
        return
    }

    existingServiceInfo.ServiceInfo1 = updatedServiceInfo.ServiceInfo1
    existingServiceInfo.ServiceInfo2 = updatedServiceInfo.ServiceInfo2
    existingServiceInfo.Tags = updatedServiceInfo.Tags
    existingServiceInfo.AdditionalInfo = updatedServiceInfo.AdditionalInfo

    if err := config.DB.Save(&existingServiceInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar informações do serviço"})
        return
    }

    c.JSON(http.StatusOK, existingServiceInfo)
}

func DeleteServiceInfo(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Delete(&models.ServiceInfo{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir informações do serviço"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Informações do serviço deletadas"})
}
