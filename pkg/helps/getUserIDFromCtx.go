package helps

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUserIDFromCtx(ctx *gin.Context) (int, error) {
	userID, ok := ctx.Get("user_id")
	if !ok {
		return 0, errors.New("error with get user_id from ctx")
	}

	useridInt, ok := userID.(int)
	if !ok {
		return 0, errors.New("error with convert user_id from ctx")
	}

	return useridInt, nil
}
