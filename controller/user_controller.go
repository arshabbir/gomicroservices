package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
	"strconv"
)

func GetUserController(resp http.ResponseWriter, req *http.Request) {

	value := req.URL.Query().Get("id")
	//userid, _ := strconv.ParseInt(value, 10, 32)
	userid, _ := strconv.Atoi(value)
	resp.Header().Set("content-type", "application/json")

	var user model.User

	user = service.GetUser(userid)
	json.NewEncoder(resp).Encode(&user)
}
