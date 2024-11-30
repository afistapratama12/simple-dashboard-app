package wrapper

import "github.com/gin-gonic/gin"

type ResponseAPI struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, status int, data interface{}) {
	var message string

	if status >= 400 && status <= 599 {
		message = "failed"
	} else {
		message = "success"
	}

	var resp = ResponseAPI{
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(status, resp)
}

func ResponseJSONWithMessage(c *gin.Context, status int, message string) {
	var data = map[string]interface{}{
		"message": message,
	}

	ResponseJSON(c, status, data)
}

func ResponseError(c *gin.Context, status int, err error) {
	ResponseJSON(c, status, ErrorData(err.Error()))
}

func ErrorData(err string) map[string]interface{} {
	return map[string]interface{}{
		"error": err,
	}
}
