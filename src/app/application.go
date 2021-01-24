package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uuthman/bookstore_oauth-api/src/domain/access_token"
	"github.com/uuthman/bookstore_oauth-api/src/http"
	"github.com/uuthman/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)
func StartApplication(){
	fmt.Println("Starting")
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	if err := router.Run(":8080"); err != nil{
		panic(err)
	}
}