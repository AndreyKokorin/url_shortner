package routers

import (
	_ "URL_shortner/docs"
	handler_Links "URL_shortner/internal/handler/links"
	handler "URL_shortner/internal/handler/user"
	"URL_shortner/internal/middlewares"
	linkRepository "URL_shortner/internal/repository/link_repository"
	userRepository "URL_shortner/internal/repository/user"
	"URL_shortner/internal/service/links"
	service "URL_shortner/internal/service/user"
	"URL_shortner/pkg/helps"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(engine *gin.Engine, db *sql.DB) {
	helps.Cors(engine)

	api := engine.Group("/api")

	postgresUserRep := userRepository.NewPostgresUserRepository(db)
	userService := service.NewUserService(postgresUserRep)
	userHandlers := handler.NewUserHandler(userService)

	auth := api.Group("/auth")
	{
		auth.POST("/register", userHandlers.RegisterUserHandler)
		auth.POST("/login", userHandlers.LogInHandler)
	}

	postgresLinksRepository := linkRepository.NewPostgresLinkRepository(db)
	linksService := links.NewLinkService(postgresLinksRepository)
	linksHandlers := handler_Links.NewLinkHandler(linksService)

	short := api.Group("/short")
	{
		short.POST("/new-link", middlewares.AuthMiddleware(linksHandlers.ShortenNewLinkHandler))
		short.DELETE("/delete-link/:linkID", middlewares.AuthMiddleware(linksHandlers.DeleteLinkHandler))
		short.GET("/links", middlewares.AuthMiddleware(linksHandlers.GetUserLinksHandler))
	}

	engine.GET("/:short", linksHandlers.RedirectHandler)
}
