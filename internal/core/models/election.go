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

type ElectionResultResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Dob        string `json:"dob"`
	BioLink    string `json:"bioLink"`
	ImageLink  string `json:"imageLink"`
	Policy     string `json:"policy"`
	VotedCount int    `json:"votedCount"`
	Percentage string `json:"percentage"`
}
