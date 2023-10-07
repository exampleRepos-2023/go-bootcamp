package models

import (
	"github.com/google/uuid"
)

type Question struct {
	QuestionUuid uuid.UUID `json:"question_uuid"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt int64 `json:"created_at"`
}

type Answer struct {
	AnswerUuid uuid.UUID `json:"answer_uuid"`
	Content string `json:"content"`
	CreatedAt int64 `json:"created_at"`
}

type QuestionAnswers struct {
	QuestionUuid uuid.UUID `json:"question_uuid"`
	AnswerUuid uuid.UUID `json:"answer_uuid"`
}
