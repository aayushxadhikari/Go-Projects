package routes

import(
	"github.com/gin-gonic/gin"
	controllers "golang-jwt-project/controllers"
	"golang-jwt-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.GET("/users/:user_id", controllers.GetUser)
}