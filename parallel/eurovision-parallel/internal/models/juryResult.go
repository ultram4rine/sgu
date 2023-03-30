package models

type JuryResult struct {
	Year       string `json:"year" db:"year"`
	Contestant string `json:"contestant" db:"contestant"`
	Country    string `json:"country" db:"country"`
	Score      int    `json:"score" db:"score"`
}
