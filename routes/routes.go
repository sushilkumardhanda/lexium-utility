package routes

import (
	"lexium-utility/controllers"
	"lexium-utility/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutesPublic(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
	r.POST("/loginConfirm", controllers.LoginConfirm)

}
func RoutesProtected(r *gin.RouterGroup) {
	r.Use(middlewares.JwtAuthMiddleware())
	r.POST("/verify", controllers.Verify)
	r.POST("/logout", controllers.Logout)
	r.POST("/getSchemaList", controllers.GetScheamList)
	r.POST("/getScheduleTree",controllers.GetScheduleTree)
	r.GET("/getElement",controllers.GetElement)
}
