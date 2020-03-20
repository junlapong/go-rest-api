package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/junlapong/go-rest-api/api/auth"
	"github.com/junlapong/go-rest-api/api/model"
	"github.com/junlapong/go-rest-api/api/response"
	"github.com/junlapong/go-rest-api/api/util/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	response.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := model.User{}

	err = server.DB.Debug().Model(model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = model.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
