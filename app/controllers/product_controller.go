package controllers

import (
    "net/http"
    "go-api/app/models"
    "go-api/config"
    "github.com/gin-gonic/gin"
    "go-api/utils"
)

func GetProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := config.DB.Where("product_id = ?", id).First(&product).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")     
    pageSizeStr := c.DefaultQuery("page_size", "10") 

    page := utils.ParseInt(pageStr)
    pageSize := utils.ParseInt(pageSizeStr)

    var products []models.Product

    offset := (page - 1) * pageSize

    if err := config.DB.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
        return
    }

    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var newProduct models.Product
    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&newProduct).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var updatedProduct models.Product
    if err := c.ShouldBindJSON(&updatedProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&models.Product{}).Where("product_id = ?", id).Updates(&updatedProduct).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    updatedProduct.ProductId = id
    c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Where("product_id = ?", id).Delete(&models.Product{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted", "productId": id})
}
