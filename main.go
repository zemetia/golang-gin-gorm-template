package main

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/config"
	repository "fp-mbd-amidrive/domain/repository"
	service "fp-mbd-amidrive/domain/service"
	controller "fp-mbd-amidrive/http/controller"
	middleware "fp-mbd-amidrive/http/middleware"
	routes "fp-mbd-amidrive/http/routes"
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
	routes.UserRoutes(server, userController, jwtService)

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
