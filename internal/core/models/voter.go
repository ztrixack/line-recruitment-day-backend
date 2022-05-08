package models

type Voter struct {
	Model
	CandidateId int    `json:"candidateId"`
	NationalId  string `json:"nationalId"`
}

type CreateVoterData struct {
	CandidateId int    `json:"candidateId"`
	NationalId  string `json:"nationalId"`
}

type VoterResponse struct {
	Id          int    `json:"id"`
	CandidateId int    `json:"candidateId"`
	NationalId  string `json:"nationalId"`
}
