package controllers

import (
    "net/http"
    "go-api/app/models"
    "github.com/gin-gonic/gin"
    "go-api/config"
    "go-api/utils"
)

func GetSubscriptions(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")         
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var subscriptions []models.Subscription

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&subscriptions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar assinaturas"})
        return
    }

    c.JSON(http.StatusOK, subscriptions)
}

func GetSubscription(c *gin.Context) {
    var subscription models.Subscription
    id := c.Param("id")

    if err := config.DB.First(&subscription, "subscription_id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Assinatura não encontrada"})
        return
    }

    c.JSON(http.StatusOK, subscription)
}

func CreateSubscription(c *gin.Context) {
    var subscription models.Subscription
    if err := c.ShouldBindJSON(&subscription); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&subscription).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar assinatura"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": subscription.SubscriptionId})
}

func UpdateSubscription(c *gin.Context) {
    var subscription models.Subscription
    id := c.Param("id")

    if err := config.DB.First(&subscription, "subscription_id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Assinatura não encontrada"})
        return
    }

    if err := c.ShouldBindJSON(&subscription); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&subscription).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar assinatura"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Assinatura atualizada com sucesso"})
}

func DeleteSubscription(c *gin.Context) {
    var subscription models.Subscription
    id := c.Param("id")

    if err := config.DB.First(&subscription, "subscription_id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Assinatura não encontrada"})
        return
    }

    if err := config.DB.Delete(&subscription).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir assinatura"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Assinatura excluída com sucesso"})
}
