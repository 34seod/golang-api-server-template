package testutils

import (
	"encoding/json"
	"golang-api-server-template/server"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func ApiSetup() (*gin.Engine, *httptest.ResponseRecorder) {
	if router == nil {
		router = server.SetRouter()
	}
	w := httptest.NewRecorder()
	return router, w
}

func ToJson(body any) string {
	response, _ := json.Marshal(body)
	return string(response)
}
