package controller

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nrule/rules"
	"github.com/NubeIO/nrule/storage"
	"github.com/gin-gonic/gin"
	"time"
)

type RulesBody struct {
	Script interface{} `json:"script"`
	Name   string      `json:"name"`
}

//func (inst *Controller) Dry2(c *gin.Context) {
//	inst.Client.Err = ""
//	inst.Client.Return = nil
//	inst.Client.TimeTaken = ""
//	start := time.Now()
//	var body *RulesBody
//	err := c.ShouldBindJSON(&body)
//	if err != nil {
//		inst.Client.Err = err.Error()
//		inst.Client.TimeTaken = time.Since(start).String()
//		reposeHandler(inst.Client, nil, c)
//		return
//	}
//
//	name := body.Name
//	if !inst.Rules.RuleExists(name) {
//		err = inst.Rules.AddRule(name, fmt.Sprint(body.Script), inst.Props)
//		if err != nil {
//			inst.Client.Err = err.Error()
//			inst.Client.TimeTaken = time.Since(start).String()
//			reposeHandler(inst.Client, nil, c)
//			return
//		}
//	}
//
//	if inst.Rules.RuleLocked(name) {
//		inst.Client.Err = fmt.Sprintf("rule: %s is already being processed", name)
//		inst.Client.TimeTaken = time.Since(start).String()
//		reposeHandler(inst.Client, nil, c)
//		return
//	}
//
//	err = inst.Rules.Execute(name)
//	if err != nil {
//		inst.Client.Err = err.Error()
//		inst.Client.TimeTaken = time.Since(start).String()
//		reposeHandler(inst.Client, nil, c)
//		return
//	}
//	err = inst.Rules.RemoveRule(name)
//	if err != nil {
//		inst.Client.Err = err.Error()
//		inst.Client.TimeTaken = time.Since(start).String()
//		reposeHandler(inst.Client, nil, c)
//		return
//	}
//	if err != nil {
//		inst.Client.Err = err.Error()
//		inst.Client.TimeTaken = time.Since(start).String()
//		reposeHandler(inst.Client, nil, c)
//	} else {
//		inst.Client.TimeTaken = time.Since(start).String()
//		pprint.PrintJSON(inst.Client)
//		reposeHandler(inst.Client, nil, c)
//	}
//
//}

func (inst *Controller) Dry(c *gin.Context) {
	start := time.Now()
	inst.Client.Err = ""
	inst.Client.Return = nil
	inst.Client.TimeTaken = ""

	var body *rules.Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		inst.Client.Err = err.Error()
		reposeHandler(inst.Client, nil, c)
		return
	}

	name := uuid.ShortUUID("")
	schedule := "1 sec"
	script := fmt.Sprint(body.Script)

	newRule := &rules.AddRule{
		Name:     name,
		Script:   script,
		Schedule: schedule,
		Props:    inst.Props,
	}
	err = inst.Rules.AddRule(newRule)
	if err != nil {
		inst.Client.Err = err.Error()
		reposeHandler(inst.Client, nil, c)
		return
	}
	value, err := inst.Rules.ExecuteAndRemove(name, inst.Props, false)
	if err != nil {
		inst.Client.Err = err.Error()
		reposeHandler(inst.Client, nil, c)
		return
	}
	inst.Client.Return = value
	inst.Client.TimeTaken = time.Since(start).String()
	reposeHandler(inst.Client, nil, c)
}

func (inst *Controller) RunExisting(c *gin.Context) {

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
