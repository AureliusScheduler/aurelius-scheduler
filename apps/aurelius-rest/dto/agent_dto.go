package dto

type AgentDto struct {
	ID           string `json:"id"`
	Service      string `json:"service"`
	AgentStack   string `json:"agent_stack"`
	AgentVersion string `json:"agent_version"`
	RegisteredAt string `json:"registered_at"`
	LastSeenAt   string `json:"last_seen_at"`
}
