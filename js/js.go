package js

import (
	"errors"
	"github.com/NubeIO/nrule"
	"github.com/dop251/goja"
)

type Rule struct {
	vm *goja.Runtime
	js string
}

type RuleMap map[string]*Rule

type RuleEngine struct {
	rules RuleMap
}

func NewRuleEngine() *RuleEngine {
	re := &RuleEngine{rules: RuleMap{}}
	return re
}

func (r *RuleEngine) Start() error {
	return nil
}

func (r *RuleEngine) Stop() error {
	return nil
}

func (r *RuleEngine) AddRule(name, script string, props nrule.PropertiesMap) error {
	_, ok := r.rules[name]
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
	r.rules[name] = &rule
	return nil
}

func (r *RuleEngine) RemoveRule(name string) error {
	delete(r.rules, name)
	return nil
}

func (r *RuleEngine) RuleCount() int {
	return len(r.rules)
}

func (r *RuleEngine) Execute(name string) error {
	rule, ok := r.rules[name]
	if !ok {
		return errors.New("rule does not exist")
	}
	_, err := rule.vm.RunString(rule.js)
	return err
}

func (r *RuleEngine) ModifyRule(name, script string) error {
	rule, ok := r.rules[name]
	if !ok {
		return errors.New("rule does not exist")
	}
	rule.js = script
	return nil
}
