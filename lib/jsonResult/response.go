package jsonResult

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result.
type Result struct {
	Code int         `json:"code"` // return code, 0 for succ
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "ok",
		Data: nil,
	}
}

func CreateSuccess(c *gin.Context, data interface{}) {
	json := NewResult()
	json.Data = data
	c.JSON(
		http.StatusOK,
		json,
	)
	c.Abort()
}

func CreateNotContent(c *gin.Context) {
	json := NewResult()
	json.Data = struct{}{}
	json.Msg = "No More Content"
	json.Code = 204

	c.JSON(
		http.StatusOK,
		json,
	)
	c.Abort()
}

func CreateError(c *gin.Context) {
	json := NewResult()
	json.Data = struct{}{}
	json.Msg = "fail"
	json.Code = -1

	c.JSON(
		http.StatusOK,
		json,
	)
	c.Abort()
}

func CreateErrorWithMsg(c *gin.Context, msg string) {
	json := NewResult()
	json.Data = make(map[string]interface{})
	json.Msg = msg
	json.Code = -1

	c.JSON(
		http.StatusOK,
		json,
	)
	c.Abort()
}

func CreateErrorParams(c *gin.Context) {
	json := NewResult()
	json.Data = struct{}{}
	json.Msg = "params error"
	json.Code = -1

	c.JSON(
		http.StatusOK,
		json,
	)
	c.Abort()
}

func CreateSuccessByList(c *gin.Context, total interface{}, content interface{}) {
	json := NewResult()
	json.Data = gin.H{
		"total":   total,
		"content": content,
	}
	c.JSON(
		http.StatusOK,
		json,
	)
	c.Abort()
}

func CreateErrorRequest(c *gin.Context) {
	json := NewResult()
	json.Data = struct{}{}
	json.Msg = "The request is frequent"
	json.Code = -1
	c.JSON(
		http.StatusOK,
		json,
	)
	c.Abort()
}
