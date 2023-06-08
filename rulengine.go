package nrule

type PropertiesMap map[string]interface{}
type FunctionsMap map[string]interface{}

type RuleEngine interface {
	Start() error
	Stop() error
	Execute(name string) error

	// AddRule
	// @param [in] name
	// @param [in] script
	// @param [in] properties
	AddRule(name, script string, properties PropertiesMap) error

	// RemoveRule
	RemoveRule(name string) error

	// RulesCount
	RulesCount() int

	// ModifyRule
	// @param [in] name
	// @param [in] script
	ModifyRule(name, script string) error
}
