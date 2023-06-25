package storage

type RQLRule struct {
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	LatestRunDate string `json:"latest_run_date"`
	Script        string `json:"script"`
	Schedule      string `json:"schedule"`
	Enable        bool   `json:"enable"`
}

type RQLVariables struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Variable any    `json:"variable"`
	Password string `json:"password"`
}
