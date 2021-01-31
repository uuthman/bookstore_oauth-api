 package rest

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"testing"

// 	"github.com/federicoleon/go-httpclient/gohttp"
// 	"github.com/stretchr/testify/assert"
// )

// func TestMain(m *testing.M){
// 	gohttp.StartMockServer()
// 	os.Exit(m.Run())
// }

// func TestLoginUserTimeoutFromApi(t *testing.T){
// 	gohttp.FlushMocks()
// 	gohttp.AddMock(
// 		gohttp.Mock{
// 			Url: "https://api.bookstore.com/users/login",
// 			Method: http.MethodPost,
// 			RequestBody: `{"email":"email@gmail.com","password":"the-pass"}`,
// 			ResponseStatusCode: -1,
// 			ResponseBody: `{}`,
// 		},
// 	)

// 	repository := usersRepository{}

// 	user, err := repository.LoginUser("email@gmail.com","the-pass")
// 	fmt.Println(user)
// 	fmt.Println(err)
// 	// assert.Nil(t,user)
// 	// assert.NotNil(t,err)
// 	// assert.EqualValues(t,http.StatusInternalServerError, err.Status)
// 	// assert.EqualValues(t,"invalid restclient response when trying to login user", err.Message)
// }

// func TestLoginUserInvalidErrorInterface(t *testing.T){


// 	repository := usersRepository{}

// 	user, err := repository.LoginUser("email@gmail.com","the-pass")
// 	assert.Nil(t,user)
// 	assert.NotNil(t,err)
// 	assert.EqualValues(t,http.StatusInternalServerError, err.Status)
// 	assert.EqualValues(t,"invalid restclient response when trying to login user", err.Message)
// }

// func TestLoginUserInvalidLoginCredentials(t *testing.T){
	

// 	repository := usersRepository{}

// 	user, err := repository.LoginUser("email@gmail.com","the-pass")
// 	assert.Nil(t,user)
// 	assert.NotNil(t,err)
// 	assert.EqualValues(t,http.StatusNotFound, err.Status)
// 	assert.EqualValues(t,"invalid login credentials", err.Message)


// }

// func TestLoginUserInvalidUserJsonResponse(t *testing.T){

// }

// func TestLoginUserNoError(t *testing.T){

// }