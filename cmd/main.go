package main

import (
	"URL_shortner/config"
	"URL_shortner/database"
	routers "URL_shortner/internal/router"
	"URL_shortner/pkg/helps"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Сокращатель ссылок
// @version 1.0
// @description API предназначена для сокращения URL-адресов и мониторинга статистики переходов по ним.
// @host localhost:8080
// @BasePath /api/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите токен в формате "Bearer <token>"
func main() {
	cfg := config.LoadConfig()

	db, err := database.PostgresDBInit(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения бд - " + err.Error())
	}

	router := gin.Default()
	helps.Cors(router)
	routers.SetupUserRouter(router, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(cfg.LOCAL_PORT)
}
