package controller

import (
	"fmt"

	"github.com/NubeIO/nrule/apirules"
	"github.com/NubeIO/nrule/rules"
	"github.com/NubeIO/nrule/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Rules   *rules.RuleEngine
	Client  *apirules.Client
	Props   rules.PropertiesMap
	Storage storage.IStorage
}

type Response struct {
	StatusCode   int         `json:"status_code"`
	ErrorMessage string      `json:"error_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func reposeHandler(body interface{}, err error, c *gin.Context, statusCode ...int) {
	var code int
	if err != nil {
		if len(statusCode) > 0 {
			code = statusCode[0]
		} else {
			code = http.StatusNotFound
		}
		msg := Message{
			Message: fmt.Sprintf("flow: %s", err.Error()),
		}
		c.JSON(code, msg)
	} else {
		if len(statusCode) > 0 {
			code = statusCode[0]
		} else {
			code = http.StatusOK
		}
		c.JSON(code, body)
	}
}

type Message struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (inst *Controller) Ping(c *gin.Context) {
	reposeHandler("hello", nil, c)
}
