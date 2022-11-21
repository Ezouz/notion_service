package models

type ProcessDB struct {
	ID     string   `json:"db,omitempty"`
	Check  []string `json:"check,omitempty"`
	Notify string   `json:"notify,omitempty"`
}
