package managers

import (
	"errors"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/persistence"
	"time"
)

var QM QuestionManager = QuestionManagerImpl{
	QuestionDao: persistence.Questions,
	AnswerDao: persistence.Answers,
	QuestionAnswersDao: persistence.QuestionAnswersS,
}

type QuestionManager interface {
	CreateQuestion(request models.CreateQuestionRequest) (models.CreateQuestionResponse, error)
	DeleteQuestion(request models.DeleteQuestionRequest) (models.DeleteQuestionResponse, error)
	GetQuestions() models.GetQuestionsResponse
}

type QuestionManagerImpl struct {
	QuestionDao persistence.QuestionDao
	AnswerDao persistence.AnswerDao
	QuestionAnswersDao persistence.QuestionAnswersDao
}

func (q QuestionManagerImpl) CreateQuestion(request models.CreateQuestionRequest) (models.CreateQuestionResponse, error) {
	question := models.Question{
		QuestionUuid: uuid.New(),
		Title: request.Title,
		Description: request.Description,
		CreatedAt: time.Now().UnixMilli(),
	}
	return models.CreateQuestionResponse{Question: question}, q.QuestionDao.UpsertQuestion(question)
}

func (q QuestionManagerImpl) DeleteQuestion(request models.DeleteQuestionRequest) (models.DeleteQuestionResponse, error)  {
	questionAnswersS := q.QuestionAnswersDao.SelectQuestionAnswersByQuestionUuid(request.QuestionUuid)

	var answerUuids []uuid.UUID
	for _, questionAnswers := range questionAnswersS {
		answerUuids = append(answerUuids, questionAnswers.AnswerUuid)
	}
	if err := q.AnswerDao.DeleteAnswers(answerUuids); err != nil {
		return models.DeleteQuestionResponse{}, errors.New("unable to delete answers")
	}

	if err := q.QuestionDao.DeleteQuestion(request.QuestionUuid); err != nil {
		return models.DeleteQuestionResponse{}, errors.New("unable to delete question")
	}

	if err := q.QuestionAnswersDao.DeleteQuestionAnswersByQuestionUuid(request.QuestionUuid); err != nil {
		return models.DeleteQuestionResponse{}, errors.New("unable to delete questionAnswers")
	}

	return models.DeleteQuestionResponse{}, nil
}

func (q QuestionManagerImpl) GetQuestions() models.GetQuestionsResponse {
	questions := q.QuestionDao.SelectQuestions()
	response := models.GetQuestionsResponse{
		Questions: questions,
	}
	return response
}