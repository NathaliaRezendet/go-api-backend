package main

import (
    "github.com/gin-gonic/gin"
    "go-api/config"
    "go-api/routes"
)

func main() {
    config.InitDatabase()

    server := gin.Default()

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
