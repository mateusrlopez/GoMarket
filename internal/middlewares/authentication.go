package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mateusrlopez/go-market/internal/services"
)

func AuthenticationMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")

		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing authorization header"})
			return
		}

		token := strings.Split(header, " ")[1]

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing token in header"})
			return
		}

		user, err := authService.ValidateToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
