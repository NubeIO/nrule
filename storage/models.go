package storage

type RQLRule struct {
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	LatestRunDate string `json:"latest_run_date"`
	Script        string `json:"script"`
}
