package models

type TargetsList []struct {
	Targets []string `json:"targets"`
	Labels  struct {
		MetricsPath string `json:"__metrics_path__"`
	} `json:"labels"`
}

type Single struct {
	ENVIRONMENT string
	APP_VERSION string
	LOG_LEVEL   string
}
