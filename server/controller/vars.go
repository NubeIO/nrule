package controller

import (
	"github.com/NubeIO/nrule/storage"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) SelectAllVariables(c *gin.Context) {
	resp, err := inst.Storage.SelectAllVariables()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) SelectVariable(c *gin.Context) {
	resp, err := inst.Storage.SelectVariable(c.Param("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) AddVariable(c *gin.Context) {
	var body *storage.RQLVariables
	err := c.ShouldBindJSON(&body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	resp, err := inst.Storage.AddVariable(body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) UpdateVariable(c *gin.Context) {
	var body *storage.RQLVariables
	err := c.ShouldBindJSON(&body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	resp, err := inst.Storage.UpdateVariable(c.Param("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) DeleteVariable(c *gin.Context) {
	err := inst.Storage.DeleteVariable(c.Param("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "ok"}, err, c)
}
