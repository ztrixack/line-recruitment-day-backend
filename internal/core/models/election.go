package models

type Election struct {
	Model
	Enable bool   `json:"enable"`
	State  string `json:"state"`
}

type UpdateElectionData struct {
	Enable bool   `json:"enable"`
	State  string `json:"state"`
}

type ElectionResponse struct {
	Status string `json:"status"`
	Enable bool   `json:"enable"`
	State  string `json:"state"`
}
