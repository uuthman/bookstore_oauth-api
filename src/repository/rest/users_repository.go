package rest

import (
	"encoding/json"
	"time"

	"github.com/federicoleon/go-httpclient/gohttp"
	"github.com/uuthman/bookstore_oauth-api/src/domain/users"
	"github.com/uuthman/bookstore_oauth-api/src/utils/errors"
)

var (
	client = gohttp.NewBuilder().SetConnectionTimeout(100 * time.Millisecond).Build()
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response,err := client.Post("https://api.bookstore.com/users/login", request)


	if err != nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}


	if response.StatusCode() > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error interface when trying to unmarshal users response")
	}

	return &user, nil

}
