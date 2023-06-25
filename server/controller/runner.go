package controller

import (
	"fmt"
	"github.com/NubeIO/nrule/rules"
	"github.com/NubeIO/nrule/storage"
	"github.com/labstack/gommon/log"
	"time"
)

func (inst *Controller) addAll(allRules []storage.RQLRule) {
	for _, rule := range allRules {
		name := rule.Name
		schedule := rule.Schedule
		schedule = "10 sec"
		script := fmt.Sprint(rule.Script)

		newRule := &rules.AddRule{
			Name:     name,
			Script:   script,
			Schedule: schedule,
			Props:    inst.Props,
		}
		err := inst.Rules.AddRule(newRule)
		if err != nil {
			log.Info(fmt.Sprintf("%s", err.Error()))
		}
	}
}

func (inst *Controller) Loop() {
	var firstLoop = true
	for {
		allRules, err := inst.Storage.SelectAllRules()
		if err != nil {
			//return
		}
		if firstLoop {
			inst.addAll(allRules) // add all existing rules from DB
		}

		for _, rule := range allRules {
			//fmt.Println("rule loop name: ", rule.Name)
			canRun, err := inst.Rules.CanExecute(rule.Name)
			if err != nil {
				//fmt.Println(err)
			}
			if canRun != nil && rule.Enable {
				if canRun.CanRun {
					execute, err := inst.Rules.Execute(rule.Name, inst.Props)
					if err != nil {
						fmt.Println("RAN RULE ERR", err)
						//return
					}
					fmt.Println("RAN RULE", execute)
				}
			} else {

			}

		}
		firstLoop = false
		time.Sleep(5 * time.Second)
	}

}
