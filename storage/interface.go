package storage

type IStorage interface {
	AddRule(rc *RQLRule) (*RQLRule, error)
	UpdateRule(uuid string, rc *RQLRule) (*RQLRule, error)
	DeleteRule(uuid string) error
	SelectRule(uuid string) (*RQLRule, error)
	SelectAllRules() ([]RQLRule, error)
}
