package main

import (
	"fp-mbd-amidrive/config"
	seeder "fp-mbd-amidrive/database/seeders"
	"net/http"

	"fp-mbd-amidrive/common"

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

	var db *gorm.DB = config.SetupDatabaseConnection()

	/* Seeder Call Function Section */
	seeder.UserSeeder(db)

}
