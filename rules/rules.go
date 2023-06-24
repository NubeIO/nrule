package rules

import (
	"errors"
	"github.com/dop251/goja"
)

type PropertiesMap map[string]interface{}

type Rule struct {
	vm   *goja.Runtime
	js   string
	lock bool
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

func (inst *RuleEngine) AddRule(name, script string, props PropertiesMap) error {
	_, ok := inst.rules[name]
	if ok {
		return errors.New("rule logic already exists")
	}

	vm := goja.New()
	if vm == nil {
		return errors.New("create js vm failed")
	}

	for k, v := range props {
		err := vm.Set(k, v)
		if err != nil {
			return err
		}
	}

	var rule Rule
	rule.vm = vm
	rule.js = script
	inst.rules[name] = &rule
	return nil
}

func (inst *RuleEngine) RemoveRule(name string) error {
	delete(inst.rules, name)
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
	rule, _ := inst.rules[name]
	return rule.lock
}

func (inst *RuleEngine) Execute(name string) error {
	rule, ok := inst.rules[name]
	rule.lock = true
	if !ok {
		return errors.New("rule does not exist")
	}
	_, err := rule.vm.RunString(rule.js)
	rule.lock = false
	return err
}

func (inst *RuleEngine) ModifyRule(name, script string) error {
	rule, ok := inst.rules[name]
	if !ok {
		return errors.New("rule does not exist")
	}
	rule.js = script
	return nil
}
