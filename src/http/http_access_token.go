package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uuthman/bookstore_oauth-api/src/services/access_token"
	"github.com/uuthman/bookstore_oauth-api/src/utils/errors"
	atDomain "github.com/uuthman/bookstore_oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface{
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct{
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler{
	return &accessTokenHandler{
		service: service,
	}
}


func (handler *accessTokenHandler) GetById(c *gin.Context){
	accessToken,err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context){
	var at atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := handler.service.Create(at)

	if err != nil{
		c.JSON(err.Status, err)
		return 	
	}

	c.JSON(http.StatusCreated, user)
}