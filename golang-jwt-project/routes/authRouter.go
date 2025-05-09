package routes

import (
	"golang-jwt-project/controllers" 
	"github.com/gin-gonic/gin"
)

// Corrected function signature
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
