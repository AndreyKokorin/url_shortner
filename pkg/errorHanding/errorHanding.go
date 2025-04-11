package pkg

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func ResponseAndLogError(ctx *gin.Context, err error, status int, massage string) {
	slog.Error(err.Error())
	ctx.JSON(status, gin.H{"error": massage})
}
