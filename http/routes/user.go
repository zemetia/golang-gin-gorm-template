package routes

import (
	"golang-gin-gorm-template/domain/repository"
	service "golang-gin-gorm-template/domain/service"
	controller "golang-gin-gorm-template/http/controller"
	middleware "golang-gin-gorm-template/http/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, UserController controller.UserController, userRepository repository.UserRepository, jwtService service.JWTService) {
	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("", UserController.RegisterUser)
		userRoutes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
		userRoutes.POST("/login", UserController.LoginUser)
		// Jika ada Batasan Role nya
		userRoutes.DELETE("/", middleware.Authenticate(jwtService), middleware.Role(userRepository, []string{"admin"}), UserController.DeleteUser)
		userRoutes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
		userRoutes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
	}
}
