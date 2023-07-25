package main

import (
	"golang-gin-gorm-template/common"
	"golang-gin-gorm-template/config"
	repository "golang-gin-gorm-template/domain/repository"
	service "golang-gin-gorm-template/domain/service"
	controller "golang-gin-gorm-template/http/controller"
	middleware "golang-gin-gorm-template/http/middleware"
	routes "golang-gin-gorm-template/http/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		res := common.BuildErrorResponse("Gagal Terhubung ke Server", err.Error(), common.EmptyObj{})
		(*gin.Context).JSON((&gin.Context{}), http.StatusBadGateway, res)
		return
	}

	var (
		db *gorm.DB = config.SetupDatabaseConnection()

		jwtService service.JWTService = service.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.UserRoutes(server, userController, userRepository, jwtService)

	port := os.Getenv("PORT")
	ip := os.Getenv("IP")
	if port == "" {
		port = "8000"
	}
	if ip == "" {
		ip = "localhost"
	}
	server.Run(ip + ":" + port)
}
