package db

import (
	"github.com/uuthman/bookstore_oauth-api/src/clients/cassandra"
	"github.com/uuthman/bookstore_oauth-api/src/domain/access_token"
	"github.com/uuthman/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository{
	return &dbRepository{}
}

type DbRepository interface{
	GetById(string) (*access_token.AccessToken, *errors.RestErr)

}

type dbRepository struct{
	
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr){
	_,err := cassandra.GetSession()
	if err != nil {
		panic(err.Error())
	}
	return nil,errors.NewInternalServerError("database connection not implemented yet")
}
