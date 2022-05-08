package models

type Candidate struct {
	Model
	Name      string `json:"name"`
	Dob       string `json:"dob"`
	BioLink   string `json:"bioLink"`
	ImageLink string `json:"imageLink"`
	Policy    string `json:"policy"`
}

type CandidateResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Dob        string `json:"dob"`
	BioLink    string `json:"bioLink"`
	ImageLink  string `json:"imageLink"`
	Policy     string `json:"policy"`
	VotedCount int    `json:"votedCount"`
}
