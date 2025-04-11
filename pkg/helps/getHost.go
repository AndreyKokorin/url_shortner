package helps

import "github.com/gin-gonic/gin"

func GetHost(ctx *gin.Context) string {
	host := ctx.Request.Host
	scheme := "http"

	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	serverDomain := scheme + "://" + host + "/"
	return serverDomain
}
