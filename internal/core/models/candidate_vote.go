package models

type CandidateVote struct {
	CandidateId int `json:"candidateId" gorm:"unique,primaryKey;autoIncrement:false"`
	VotedCount  int `json:"votedCount"`
}

type CreateCandidateVoteData struct {
	CandidateId int `json:"candidateId"`
}

type UpdateCandidateVoteData struct {
	VotedCount int `json:"votedCount"`
}

type CandidateVoteResponse struct {
	Id         string `json:"id"`
	VotedCount int    `json:"votedCount"`
}
