package handler

import (
	"URL_shortner/internal/model"
	pkg "URL_shortner/pkg/errorHanding"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type UserHandler struct {
	userService model.UserService
}

func NewUserHandler(userService model.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// RegisterUserHandler godoc
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя по email и паролю. Email должен быть уникальным
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "Данные пользователя (email и пароль)"
// @Success 201 {object} map[string]string "Пользователь успешно создан"
// @Failure 400 {object} map[string]string "Ошибка валидации входных данных"
// @Failure 409 {object} map[string]string "Пользователь с таким email уже существует"
// @Router /auth/register [post]
func (userHandler UserHandler) RegisterUserHandler(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusBadRequest, "error with get data")
		return
	}

	err = userHandler.userService.RegisterUser(&user)
	if err != nil {
		//исключаем дублирование email
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			pkg.ResponseAndLogError(ctx, err, http.StatusConflict, "user with this email already exists")
			return
		}
		pkg.ResponseAndLogError(ctx, err, http.StatusConflict, "user with register user")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user create success"})
}
