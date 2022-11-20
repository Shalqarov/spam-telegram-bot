package web

import (
	"encoding/json"
	"net/http"
)

func statusMessage(status int) []byte {
	resp := make(map[string]string)
	switch status {
	case http.StatusOK:
		resp["message"] = "OK"
		resp["code"] = "200"
	case http.StatusBadRequest:
		resp["message"] = "Bad Request"
		resp["code"] = "400"
	case http.StatusNotFound:
		resp["message"] = "Not Found"
		resp["code"] = "404"
	case http.StatusMethodNotAllowed:
		resp["message"] = "Method Not Allowed"
		resp["code"] = "405"
	default:
		resp["message"] = "Internal Server Error"
		resp["code"] = "500"
	}
	jsonResp, _ := json.Marshal(resp)
	return jsonResp
}
