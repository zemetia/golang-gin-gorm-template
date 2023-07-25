package middleware

import (
	"fp-mbd-amidrive/common"
	service "fp-mbd-amidrive/domain/service"
	"fp-mbd-amidrive/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Role(userService service.UserService, role []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get("userID")
		user, err := userService.FindUserByID(ctx.Request.Context(), userID.(string))

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
