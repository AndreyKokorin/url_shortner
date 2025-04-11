package handler_Links

import (
	pkg "URL_shortner/pkg/errorHanding"
	"URL_shortner/pkg/helps"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// GetUserLinksHandler godoc
// @Summary Получить все ссылки пользователя
// @Description Возвращает список всех сокращённых ссылок, созданных авторизованным пользователем
// @Tags links
// @Produce json
// @Success 200 {array} model.Link "Список ссылок"
// @Failure 400 {object} map[string]string "Ошибка получения user_id или ссылок"
// @Router /short/links [get]
// @Security BearerAuth
func (lh *LinkHandler) GetUserLinksHandler(ctx *gin.Context) {
	userID, err := helps.GetUserIDFromCtx(ctx)
	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusBadRequest, "error with get your user_id")
		return
	}

	links, err := lh.linkService.GetUserLinks(userID)
	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusBadRequest, "error with get your user_id")
		return
	}

	for i := range links {
		links[i].ShortUrl = helps.GetHost(ctx) + links[i].ShortUrl
		slog.Info(links[i].ShortUrl)
	}

	ctx.JSON(http.StatusOK, links)
}
