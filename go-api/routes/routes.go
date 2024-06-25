package routes

import (
    "github.com/gin-gonic/gin"
    "go-api/app/controllers"
)

func SetupRoutes(router *gin.Engine, auth gin.HandlerFunc) {
    router.GET("/client/:id", auth, controllers.GetClient)
    router.GET("/clients", auth, controllers.GetClients)
    router.POST("/client", auth, controllers.CreateClient)
    router.PUT("/client/:id", auth, controllers.UpdateClient)
    router.DELETE("/client/:id", auth, controllers.DeleteClient)

    router.GET("/benefit/:id", auth, controllers.GetBenefitEntity)
    router.GET("/benefits", auth, controllers.GetBenefitEntities)
    router.POST("/benefit", auth, controllers.CreateBenefitEntity)
    router.PUT("/benefit/:id", auth, controllers.UpdateBenefitEntity)
    router.DELETE("/benefit/:id", auth, controllers.DeleteBenefitEntity)

    router.GET("/partner/:id", auth, controllers.GetPartner)
    router.GET("/partners", auth, controllers.GetPartners)
    router.POST("/partner", auth, controllers.CreatePartner)
    router.PUT("/partner/:id", auth, controllers.UpdatePartner)
    router.DELETE("/partner/:id", auth, controllers.DeletePartner)

    router.GET("/product/:id", auth, controllers.GetProduct)
    router.GET("/products", auth, controllers.GetProducts)
    router.POST("/product", auth, controllers.CreateProduct)
    router.PUT("/product/:id", auth, controllers.UpdateProduct)
    router.DELETE("/product/:id", auth, controllers.DeleteProduct)

    router.GET("/billing/:id", auth, controllers.GetBilling)
    router.GET("/billings", auth, controllers.GetBillings)
    router.POST("/billing", auth, controllers.CreateBilling)
    router.PUT("/billing/:id", auth, controllers.UpdateBilling)
    router.DELETE("/billing/:id", auth, controllers.DeleteBilling)

    router.GET("/resource_usage/:id", auth, controllers.GetResourceUsage)
    router.GET("/resource_usages", auth, controllers.GetResourceUsages)
    router.POST("/resource_usage", auth, controllers.CreateResourceUsage)
    router.PUT("/resource_usage/:id", auth, controllers.UpdateResourceUsage)
    router.DELETE("/resource_usage/:id", auth, controllers.DeleteResourceUsage)

    router.GET("/service_info/:id", auth, controllers.GetServiceInfo)
    router.GET("/service_infos", auth, controllers.GetServiceInfos)
    router.POST("/service_info", auth, controllers.CreateServiceInfo)
    router.PUT("/service_info/:id", auth, controllers.UpdateServiceInfo)
    router.DELETE("/service_info/:id", auth, controllers.DeleteServiceInfo)

    router.GET("/subscription/:id", auth, controllers.GetSubscription)
    router.GET("/subscriptions", auth, controllers.GetSubscriptions)
    router.POST("/subscription", auth, controllers.CreateSubscription)
    router.PUT("/subscription/:id", auth, controllers.UpdateSubscription)
    router.DELETE("/subscription/:id", auth, controllers.DeleteSubscription)

    router.GET("/tag/:id", auth, controllers.GetTag)
    router.GET("/tags", auth, controllers.GetTags)
    router.POST("/tag", auth, controllers.CreateTag)
    router.PUT("/tag/:id", auth, controllers.UpdateTag)
    router.DELETE("/tag/:id", auth, controllers.DeleteTag)
}
