package dto

type JobDto struct {
	ID         string `json:"id"`
	AgentID    string `json:"agentId"`
	Name       string `json:"name"`
	Version    string `json:"version"`
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
}
