package handler_Links

import (
	pkg "URL_shortner/pkg/errorHanding"
	"URL_shortner/pkg/helps"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ErrNoPermission = errors.New("user doesn't have permission to delete the link")

// DeleteLinkHandler godoc
// @Summary Удалить ссылку пользователя
// @Description Удаляет указанную ссылку, если она принадлежит авторизованному пользователю
// @Tags links
// @Param linkID path int true "ID ссылки для удаления"
// @Success 200 {object} map[string]interface{} "Ссылка успешно удалена"
// @Failure 400 {object} map[string]string "Ошибка получения user_id или некорректный linkID"
// @Failure 500 {object} map[string]string "Ссылка не принадлежит пользователю или другая ошибка удаления"
// @Router /links/{linkID} [delete]
// @Security BearerAuth
func (lh *LinkHandler) DeleteLinkHandler(ctx *gin.Context) {
	userID, err := helps.GetUserIDFromCtx(ctx)
	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusBadRequest, "error with get your user_id")
		return
	}

	linkId, err := strconv.Atoi(ctx.Param("linkID"))
	if err != nil {
		pkg.ResponseAndLogError(ctx, err, http.StatusBadRequest, "error with convert linkID to int")
		return
	}

	err = lh.linkService.DeleteLink(linkId, userID)
	if err != nil {
		if errors.Is(err, ErrNoPermission) {
			pkg.ResponseAndLogError(ctx, err, http.StatusInternalServerError, "You are trying to delete someone else's link")
			return
		}

		pkg.ResponseAndLogError(ctx, err, http.StatusInternalServerError, "error deleting link")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Deleted_link": linkId})
}
