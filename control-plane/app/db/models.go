package db

import "gopkg.in/guregu/null.v4"

type Challenge struct {
	Model
	Name        string     `json:"name" gorm:"type:text;not null"`
	Description string     `json:"description" gorm:"type:text"`
	Filename    string     `json:"filename" gorm:"type:text;unique"`
	Caption     string     `json:"caption" gorm:"type:text"`
	CompletedAt null.Time  `json:"completed_at"`
	Questions   []Question `json:"questions"`
}

type Question struct {
	Model
	Question    string `json:"question" gorm:"type:text"`
	Answer      string `json:"answer" gorm:"type:text"`
	ChallengeID uint   `json:"challenge_id"`
}
