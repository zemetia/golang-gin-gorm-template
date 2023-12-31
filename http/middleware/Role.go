package middleware

import (
	"golang-gin-gorm-template/common"
	"golang-gin-gorm-template/domain/repository"
	"golang-gin-gorm-template/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Role(userRepository repository.UserRepository, role []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get("userID")
		user, err := userRepository.FindUserByID(ctx.Request.Context(), userID.(string))

		if err != nil {
			response := common.BuildErrorResponse("Gagal Mendapatkan User", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !helpers.Contains(role, user.Role) {
			response := common.BuildErrorResponse("Role tidak sesuai", "Tidak ada user", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Next()
	}
}
