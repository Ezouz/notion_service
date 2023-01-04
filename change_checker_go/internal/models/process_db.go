package models

type ProcessDB struct {
	ID     string            `json:"db,omitempty"`
	Check  []ProcessProperty `json:"check,omitempty"`
	Notify string            `json:"notify,omitempty"`
}

type ProcessProperty struct {
	Type   string   `json:"type,omitempty"`
	Values []string `json:"values,omitempty"`
}
