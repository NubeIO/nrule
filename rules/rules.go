package rules

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nrule/helpers/ttime"
	"github.com/dop251/goja"
	"time"
)

type PropertiesMap map[string]interface{}

type State string

const (
	Processing State = "Processing"
	Disabled   State = "Disabled"
	Completed  State = "Completed"
)

type Rule struct {
	vm                *goja.Runtime
	script            string
	lock              bool
	State             State  // processing, disabled, completed
	Schedule          string //run every 5 min, 3 hour 15 min, 3 days
	TimeCompleted     time.Time
	NextTimeScheduled time.Time
	TimeDue           string
	TimeTaken         string // 12ms

}

type RuleMap map[string]*Rule

type RuleEngine struct {
	rules  RuleMap
	Result int
}

func NewRuleEngine() *RuleEngine {
	re := &RuleEngine{rules: RuleMap{}}
	return re
}

type Body struct {
	Script   interface{} `json:"script"`
	Name     string      `json:"name"`
	Schedule string      `json:"schedule"`
}

type AddRule struct {
	Name     string
	Script   string
	Schedule string
	Props    PropertiesMap
}

func (inst *RuleEngine) AddRule(body *AddRule) error {
	name := body.Name
	script := body.Script
	props := body.Props
	sch := body.Schedule
	if inst.RuleLocked(name) {
		return errors.New(fmt.Sprintf("rule:%s is already being processed", name))
	}
	_, ok := inst.rules[name]
	if ok {
		return errors.New("rule logic already exists")
	}
	var vm *goja.Runtime
	vm = goja.New()
	if vm == nil {
		return errors.New("create script vm failed")
	}

	for k, v := range props {
		err := vm.Set(k, v)
		if err != nil {
			return err
		}
	}
	var rule Rule
	rule.vm = vm
	rule.script = script
	rule.Schedule = sch
	inst.rules[name] = &rule
	return nil
}

func (inst *RuleEngine) GetRules() (RuleMap, error) {
	return inst.rules, nil
}

func (inst *RuleEngine) GetRule(name string) (*Rule, error) {
	rule, ok := inst.rules[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("rule:%s does not exist", name))
	}
	return rule, nil
}

func (inst *RuleEngine) RemoveRule(name string) error {
	delete(inst.rules, name)
	return nil
}

// resetRule delete the VM of goja
func (inst *RuleEngine) resetRule(name string, props PropertiesMap) error {
	rule, ok := inst.rules[name]
	if !ok {
		return errors.New(fmt.Sprintf("rule:%s does not exist", name))
	}

	var vm *goja.Runtime
	vm = goja.New()
	if vm == nil {
		return errors.New("create script vm failed")
	}

	for k, v := range props {
		err := vm.Set(k, v)
		if err != nil {
			return err
		}
	}

	rule.vm = vm
	return nil
}

func (inst *RuleEngine) RuleCount() int {
	return len(inst.rules)
}

func (inst *RuleEngine) RuleExists(name string) bool {
	_, exists := inst.rules[name]
	return exists
}

func (inst *RuleEngine) RuleLocked(name string) bool {
	exists := inst.RuleExists(name)
	if !exists {
		return false
	}
	rule, _ := inst.rules[name]
	return rule.lock
}

type CanExecute struct {
	CanRun       bool    `json:"can_run"`
	TimeDueInMin float64 `json:"time_due_in_min"`
	TimeDue      string  `json:"time_due"`
}

func (inst *RuleEngine) CanExecute(name string) (*CanExecute, error) {
	rule, ok := inst.rules[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("rule:%s does not exist", name))
	}

	now := time.Now()
	nextTimeScheduled := rule.NextTimeScheduled
	dif := ttime.GetMinDifference(nextTimeScheduled, now)
	var canRun bool
	if dif <= 0 {
		canRun = true
	}
	out := &CanExecute{
		CanRun:       canRun,
		TimeDueInMin: dif,
		TimeDue:      ttime.TimePretty(nextTimeScheduled),
	}
	return out, nil
}

func (inst *RuleEngine) Execute(name string, props PropertiesMap) (goja.Value, error) {
	start := time.Now()
	rule, ok := inst.rules[name]
	rule.lock = true
	rule.State = Processing
	if !ok {
		return nil, errors.New(fmt.Sprintf("rule:%s does not exist", name))
	}
	v, err := rule.vm.RunString(rule.script)
	rule.lock = false
	rule.TimeTaken = time.Since(start).String()
	rule.State = Completed
	rule.TimeCompleted = time.Now()
	nextTime, err := ttime.AdjustTime(rule.TimeCompleted, rule.Schedule)
	rule.NextTimeScheduled = nextTime
	err = inst.resetRule(name, props)
	return v, err
}

func (inst *RuleEngine) ModifyRule(name, script string) error {
	rule, ok := inst.rules[name]
	if !ok {
		return errors.New(fmt.Sprintf("rule:%s does not exist", name))
	}
	rule.script = script
	return nil
}
