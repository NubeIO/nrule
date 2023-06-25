package controller

import (
	"fmt"
	"github.com/NubeIO/nrule/pprint"
	"time"
)

func (inst *Controller) Loop() {

	for {
		allRules, err := inst.Storage.SelectAllRules()
		if err != nil {
			//return
		}

		for _, rule := range allRules {
			//fmt.Println("rule loop name: ", rule.Name)
			canRun, err := inst.Rules.CanExecute(rule.Name)
			if err != nil {
				//fmt.Println(err)
			}
			if canRun != nil {
				pprint.PrintJSON(canRun)
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

		time.Sleep(5 * time.Second)
	}
}
