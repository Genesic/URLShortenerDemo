package testUtils

import (
	"URLShortenerDemo/pkg/errors"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func FireRequest(router *gin.Engine, method string, target string, request interface{}) *httptest.ResponseRecorder {
	req := genRequest(method, target, request)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func VerifySuccessRes(res *httptest.ResponseRecorder, statusCode int, resp interface{}) {
	So(res.Code, ShouldEqual, statusCode)
	if resp != nil {
		body, _ := ioutil.ReadAll(res.Body)
		err := json.Unmarshal(body, resp)
		So(err, ShouldEqual, nil)
	}
}

func VerifyErrorRes(res *httptest.ResponseRecorder, err *errors.ServiceError) {
	body, _ := ioutil.ReadAll(res.Body)
	resp := &errors.ServiceError{}
	jsonErr := json.Unmarshal(body, resp)
	So(jsonErr, ShouldEqual, nil)
	So(res.Code, ShouldEqual, err.Status)
	So(resp.Code, ShouldEqual, err.Code)
	So(resp.Status, ShouldEqual, err.Status)
}

func genRequest(method string, target string, request interface{}) *http.Request {
	var body io.Reader = nil
	if request != nil {
		jsonByte, _ := json.Marshal(request)
		body = bytes.NewReader(jsonByte)
	}
	return httptest.NewRequest(method, target, body)
}
