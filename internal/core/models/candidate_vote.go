package models

type CandidateVote struct {
	Model
	CandidateId int `json:"candidateId" gorm:"unique"`
	VotedCount  int `json:"votedCount"`
}

type CreateCandidateVoteData struct {
	CandidateId int `json:"candidateId"`
	VotedCount  int `json:"votedCount"`
}

type UpdateCandidateVoteData struct {
	VotedCount int `json:"votedCount"`
}

type CandidateVoteResponse struct {
	Id          int    `json:"id"`
	CandidateId int    `json:"candidateId"`
	VotedCount  string `json:"votedCount"`
}
