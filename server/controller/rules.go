package controller

import (
	"github.com/NubeIO/nrule/storage"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) SelectAllRules(c *gin.Context) {
	resp, err := inst.Storage.SelectAllRules()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) SelectRule(c *gin.Context) {
	resp, err := inst.Storage.SelectRule(c.Param("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) AddRule(c *gin.Context) {
	var body *storage.RQLRule
	err := c.ShouldBindJSON(&body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	resp, err := inst.Storage.AddRule(body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) UpdateRule(c *gin.Context) {
	var body *storage.RQLRule
	err := c.ShouldBindJSON(&body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	resp, err := inst.Storage.UpdateRule(c.Param("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(resp, err, c)
}

func (inst *Controller) DeleteRule(c *gin.Context) {
	err := inst.Storage.DeleteRule(c.Param("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "ok"}, err, c)
}
