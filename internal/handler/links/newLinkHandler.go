package handler_Links

import (
	"URL_shortner/internal/model"
	pkg "URL_shortner/pkg/errorHanding"
	"URL_shortner/pkg/helps"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LinkHandler struct {
	linkService model.LinkService
}

type OriginalUrl struct {
	OriginalUrl string `json:"original_url"`
}

func NewLinkHandler(linkService model.LinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
	}
}

// ShortenNewLinkHandler @Summary Создать короткую ссылку
// @Description Преобразует длинный URL в короткую ссылку для авторизованного пользователя
// @Tags links
// @Accept json
// @Produce json
// @Param original_url body OriginalUrl true "Оригинальный URL для сокращения"
// @Success 200 {object} map[string]string "Успешно созданная короткая ссылка"
// @Failure 400 {object} map[string]string "Ошибка валидации входных данных"
// @Failure 401 {object} map[string]string "Ошибка авторизации (нет user_id)"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /short/new-link [post]
// @Security BearerAuth
func (lh *LinkHandler) ShortenNewLinkHandler(ctx *gin.Context) {
	var originalUrl OriginalUrl

	if err := ctx.ShouldBindJSON(&originalUrl); err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusBadRequest, "error binding data")
		return
	}

	userID, ok := ctx.Get("user_id")
	if !ok {
		pkg.ResponseAndLogError(ctx, errors.New("error with your user_id"), http.StatusUnauthorized, "error with your user_id")
		return
	}

	newLink, err := lh.linkService.Shorten(userID.(int), originalUrl.OriginalUrl)
	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusInternalServerError, "error shortening new link")
		return
	}

	newLinkWithDomain := helps.GetHost(ctx) + newLink.ShortUrl

	ctx.JSON(http.StatusOK, gin.H{"new_Shorted_Link": newLinkWithDomain})
}
