package app

import (
	"github.com/gin-gonic/gin"
	"github.com/uuthman/bookstore_oauth-api/src/http"
	"github.com/uuthman/bookstore_oauth-api/src/repository/db"
	"github.com/uuthman/bookstore_oauth-api/src/repository/rest"
	"github.com/uuthman/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)
func StartApplication(){
	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(),db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	if err := router.Run(":8080"); err != nil{
		panic(err)
	}
}