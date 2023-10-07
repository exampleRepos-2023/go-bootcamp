package models

import "github.com/google/uuid"

type CreateQuestionRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

type CreateQuestionResponse struct {
	Question Question `json:"question"`
}

type GetQuestionsRequest struct {
}

type GetQuestionsResponse struct {
	Questions []Question `json:"questions"`
}

type DeleteQuestionRequest struct {
	QuestionUuid uuid.UUID `json:"question_uuid"`
}

type DeleteQuestionResponse struct {
}

type CreateAnswerRequest struct {
	QuestionUuid uuid.UUID `json:"question_uuid"`
	Content string `json:"content"`
}

type CreateAnswerResponse struct {
	Answer Answer `json:"answer"`
}

type GetAnswersRequest struct {
	QuestionUuid uuid.UUID `json:"question_uuid"`
}

type GetAnswersResponse struct {
	Answers []Answer `json:"answers"`
}

type DeleteAnswerRequest struct {
	AnswerUuid uuid.UUID `json:"answer_uuid"`
}

type DeleteAnswerResponse struct {
}