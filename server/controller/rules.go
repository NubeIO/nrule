package controller

import (
	"fmt"
	"github.com/NubeIO/nrule/helpers/uuid"
	"github.com/NubeIO/nrule/storage"
	"github.com/gin-gonic/gin"
	"time"
)

type RulesBody struct {
	Script interface{} `json:"script"`
	Name   string      `json:"name"`
}

func (inst *Controller) Dry(c *gin.Context) {
	inst.Client.Err = ""
	inst.Client.Return = nil
	inst.Client.TimeTaken = ""
	start := time.Now()
	var body *RulesBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
		return
	}

	name := body.Name
	if !inst.Rules.RuleExists(name) {
		err = inst.Rules.AddRule(name, fmt.Sprint(body.Script), inst.Props)
		if err != nil {
			inst.Client.Err = err.Error()
			inst.Client.TimeTaken = time.Since(start).String()
			reposeHandler(inst.Client, nil, c)
			return
		}
	}

	if inst.Rules.RuleLocked(name) {
		inst.Client.Err = fmt.Sprintf("rule: %s is already being processed", name)
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
		return
	}

	err = inst.Rules.Execute(name)
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
		return
	}
	err = inst.Rules.RemoveRule(name)
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
		return
	}
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
	} else {
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
	}

}

func (inst *Controller) RunExisting(c *gin.Context) {
	inst.Client.Err = ""
	inst.Client.Return = nil
	inst.Client.TimeTaken = ""
	start := time.Now()
	ruleUUID := c.Param("uuid")
	resp, err := inst.Storage.SelectRule(ruleUUID)
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(err, err, c)
		return
	}

	name := uuid.ShortUUID("")

	err = inst.Rules.AddRule(name, resp.Script, inst.Props)
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
		return
	}

	err = inst.Rules.Execute(name)

	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
		return
	}
	//err = inst.Rules.RemoveRule(name)
	//if err != nil {
	//	inst.Client.Err = err.Error()
	//	inst.Client.TimeTaken = time.Since(start).String()
	//	reposeHandler(inst.Client, nil, c)
	//	return
	//}
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
	} else {
		resp.LatestRunDate = time.Now().Format(time.RFC822)
		resp, err = inst.Storage.UpdateRule(ruleUUID, resp)
		if err != nil {
			inst.Client.Err = err.Error()
			inst.Client.TimeTaken = time.Since(start).String()
			reposeHandler(inst.Client, nil, c)
			return
		}
		inst.Client.TimeTaken = time.Since(start).String()
		reposeHandler(inst.Client, nil, c)
	}

}

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
