package handlers

import (
	"errors"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
)

var questionManagerMock = QuestionManagerMock{}
var answerManagerMock = AnswerManagerMock{}

type QuestionManagerMock struct {
	CreateQuestionRequest models.CreateQuestionRequest
	CreateQuestionResponse models.CreateQuestionResponse
	CreateQuestionError error
	CreateQuestionCalled bool

	DeleteQuestionRequest models.DeleteQuestionRequest
	DeleteQuestionResponse models.DeleteQuestionResponse
	DeleteQuestionError error
	DeleteQuestionCalled bool

	GetQuestionsResponse models.GetQuestionsResponse
	GetQuestionsCalled bool
}

func (q *QuestionManagerMock) CreateQuestion(request models.CreateQuestionRequest) (models.CreateQuestionResponse, error) {
	q.CreateQuestionCalled = true
	if q.CreateQuestionError != nil {
		return models.CreateQuestionResponse{}, q.CreateQuestionError
	}
	if request == q.CreateQuestionRequest {
		return q.CreateQuestionResponse, nil
	}
	return models.CreateQuestionResponse{}, errors.New("unable to match request")
}

func (q *QuestionManagerMock) DeleteQuestion(request models.DeleteQuestionRequest) (models.DeleteQuestionResponse, error) {
	q.DeleteQuestionCalled = true
	if request == q.DeleteQuestionRequest {
		return q.DeleteQuestionResponse, nil
	}
	return models.DeleteQuestionResponse{}, errors.New("unable to match request")
}

func (q *QuestionManagerMock) GetQuestions() models.GetQuestionsResponse {
	q.GetQuestionsCalled = true
	return q.GetQuestionsResponse
}

type AnswerManagerMock struct {
	CreateAnswerRequest models.CreateAnswerRequest
	CreateAnswerResponse models.CreateAnswerResponse
	CreateAnswerError error
	CreateAnswerCalled bool

	DeleteAnswerRequest models.DeleteAnswerRequest
	DeleteAnswerResponse models.DeleteAnswerResponse
	DeleteAnswerError error
	DeleteAnswerCalled bool

	GetAnswersRequest models.GetAnswersRequest
	GetAnswersResponse models.GetAnswersResponse
	GetAnswersCalled bool
}

func (a *AnswerManagerMock) CreateAnswer(request models.CreateAnswerRequest) (models.CreateAnswerResponse, error) {
	a.CreateAnswerCalled = true
	if a.CreateAnswerError != nil {
		return models.CreateAnswerResponse{}, a.CreateAnswerError
	}
	if request == a.CreateAnswerRequest {
		return a.CreateAnswerResponse, nil
	}
	return models.CreateAnswerResponse{}, errors.New("unable to match request")
}

func (a *AnswerManagerMock) DeleteAnswer(request models.DeleteAnswerRequest) (models.DeleteAnswerResponse, error) {
	a.DeleteAnswerCalled = true
	if a.DeleteAnswerError != nil {
		return models.DeleteAnswerResponse{}, a.CreateAnswerError
	}
	if a.DeleteAnswerError != nil {
		return models.DeleteAnswerResponse{}, a.DeleteAnswerError
	}
	if request == a.DeleteAnswerRequest {
		return a.DeleteAnswerResponse, nil
	}
	return models.DeleteAnswerResponse{}, errors.New("unable to match request")
}

func (a *AnswerManagerMock) GetAnswers(request models.GetAnswersRequest) models.GetAnswersResponse {
	a.GetAnswersCalled = true
	if request == a.GetAnswersRequest {
		return a.GetAnswersResponse
	}
	return models.GetAnswersResponse{}
}
