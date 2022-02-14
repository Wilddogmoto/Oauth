package api

import "github.com/gin-gonic/gin"

type customError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var responses = map[int]customError{

	0:  {200, "success", nil},
	2:  {400, "bad request", nil},
	3:  {400, "wrong password", nil},
	4:  {400, "name is taken", nil},
	5:  {200, "account created", nil},
	6:  {400, "invalid username", nil},
	7:  {400, "wrong password", nil},
	8:  {403, "empty authorized header", nil},
	9:  {403, "invalid authorized header", nil},
	10: {500, "Internal error", nil},
	11: {401, "signature is invalid", nil},
}

func sendResponse(id int, data interface{}, c *gin.Context) {
	c.AbortWithStatusJSON(responses[id].Code, customError{
		Code: id,
		Data: data,
		Message: responses[id].Message,
	})
	return
}
