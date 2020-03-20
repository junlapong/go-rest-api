package controller

import (
	"net/http"

	"github.com/junlapong/go-rest-api/api/response"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome to Awesome API")

}
