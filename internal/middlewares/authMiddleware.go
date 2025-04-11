package middlewares

import (
	pkg "URL_shortner/pkg/errorHanding"
	"URL_shortner/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		AuthHeader := ctx.GetHeader("Authorization")
		if AuthHeader == "" {
			pkg.ResponseAndLogError(ctx, fmt.Errorf("authorization header is missing"),
				http.StatusUnauthorized, "authorization header is missing")
			return
		}
		splitToken := strings.Split(AuthHeader, " ")

		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			pkg.ResponseAndLogError(ctx, fmt.Errorf("JWT token is wrong"), http.StatusUnauthorized, "Bearer token is missing")
			return
		}

		claims, err := jwt.ParseJWT(splitToken[1], os.Getenv("JWT_SECRET"))
		if err != nil {
			pkg.ResponseAndLogError(ctx, fmt.Errorf("JWT token is wrong"), http.StatusUnauthorized, "JWT token is wrong")
			return
		}

		ctx.Set("user_id", claims.UserID)
		next(ctx)
	}
}
