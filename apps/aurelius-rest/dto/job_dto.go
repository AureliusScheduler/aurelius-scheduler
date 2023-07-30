package dto

type JobDto struct {
	ID      string `json:"id"`
	AgentID string `json:"agentId"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Cron    string `json:"cron"`
	//NextRun    string `json:"nextRun"`
	//StartedAt  string `json:"startedAt"`
	//FinishedAt string `json:"finishedAt"`
}

type JobInput struct {
	Cron string `json:"cron"`
}
