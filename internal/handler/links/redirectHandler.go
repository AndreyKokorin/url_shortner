package handler_Links

import (
	pkg "URL_shortner/pkg/errorHanding"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// RedirectHandler godoc
// @Summary Редирект на оригинальный URL
// @Description Выполняет редирект на исходный URL, используя сокращённую ссылку
// @Tags links
// @Param short_link path string true "Сокращённая ссылка"
// @Success 302 {string} string "Редирект на оригинальный URL"
// @Failure 500 {object} map[string]string "Ошибка при обработке редиректа"
// @Router /links/{short_link} [get]
// @Security BearerAuth
func (lh *LinkHandler) RedirectHandler(ctx *gin.Context) {
	shortLink := ctx.Param("short")

	originalUrl, err := lh.linkService.GetToRedirectURL(shortLink)

	slog.Info("redirect url - " + originalUrl)

	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, originalUrl)
}
