package database

type Subject string

const (
	German      Subject = "German"
	English     Subject = "German"
	Programming Subject = "German"
)

type Grade struct {
	Subject Subject `json:"subject,omitempty"` // German
	Grade   int     `json:"grade,omitempty"`   // 1
}

func NewGrade(subject Subject, grade int) *Grade {
	return &Grade{Subject: subject, Grade: grade}
}
