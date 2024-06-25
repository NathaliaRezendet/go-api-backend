package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "go-api/config"
    "go-api/routes"
    "time"
)

func main() {
    config.InitDatabase()

    server := gin.Default()

    // Configurar CORS
    server.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    routes.SetupRoutes(server, basicAuth())

    server.Run(":8000")
}

func basicAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, pass, hasAuth := c.Request.BasicAuth()
        if !hasAuth || user != "admin" || pass != "1234" {
            c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
            c.JSON(401, gin.H{"message": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
