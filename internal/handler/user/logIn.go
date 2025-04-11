package handler

import (
	"URL_shortner/internal/model"
	pkg "URL_shortner/pkg/errorHanding"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LogInHandler godoc
// @Summary Вход пользователя
// @Description Аутентификация пользователя по email и паролю. Возвращает JWT токен при успешном входе
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "Данные пользователя (email и пароль)"
// @Success 200 {object} map[string]string "Успешная авторизация и JWT токен"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 401 {object} map[string]string "Неверный email или пароль"
// @Router /auth/login [post]
func (uh UserHandler) LogInHandler(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusBadRequest, "error with data")
		return
	}

	token, err := uh.userService.LogIn(user.Email, user.Password)
	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusUnauthorized, "wrong email or password")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
