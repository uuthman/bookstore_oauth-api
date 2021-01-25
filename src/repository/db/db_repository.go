package db

import (
	"github.com/gocql/gocql"
	"github.com/uuthman/bookstore_oauth-api/src/clients/cassandra"
	"github.com/uuthman/bookstore_oauth-api/src/domain/access_token"
	"github.com/uuthman/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token,user_id,client_id,expires FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token,user_id,client_id,expires) VALUES (?,?,?,?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires = ? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken

	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound{
			return nil, errors.NewNotFoundError("no result found with given access token")
		}
		return nil,errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr{
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken, at.AccessToken,at.UserID,at.ClientID,at.Expires).Exec(); err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr{
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,at.Expires,at.AccessToken).Exec(); err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
