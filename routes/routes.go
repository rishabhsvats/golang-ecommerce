package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishabhsvats/golang-ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/users/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.POST("/users/productview", controllers.SearchProduct())
	incomingRoutes.POST("/users/search", controllers.SearchProductByQuery())

}
