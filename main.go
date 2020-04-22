package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"

	"go-base-cleancode/infrastructure/config"
	_httpDeliver "go-base-cleancode/service/delivery/http_handler"
	// _articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	// _articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	// _authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
)

func main() {
	db := config.Init()

	route := gin.Default()

	_httpDeliver.NewHttpHandler(route, db)

	route.Run(viper.GetString("server.address"))
}
