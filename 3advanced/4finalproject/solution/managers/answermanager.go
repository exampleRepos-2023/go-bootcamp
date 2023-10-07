package managers

import (
	"errors"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/persistence"
	"log"
	"time"
)

var AM AnswerManager = AnswerManagerImpl{
	QuestionDao: persistence.Questions,
	AnswerDao: persistence.Answers,
	QuestionAnswersDao: persistence.QuestionAnswersS,
}

type AnswerManager interface {
	CreateAnswer(request models.CreateAnswerRequest) (models.CreateAnswerResponse, error)
	DeleteAnswer(request models.DeleteAnswerRequest) (models.DeleteAnswerResponse, error)
	GetAnswers(request models.GetAnswersRequest) models.GetAnswersResponse
}

type AnswerManagerImpl struct {
	QuestionDao persistence.QuestionDao
	AnswerDao persistence.AnswerDao
	QuestionAnswersDao persistence.QuestionAnswersDao
}

func (a AnswerManagerImpl) CreateAnswer(request models.CreateAnswerRequest) (models.CreateAnswerResponse, error) {
	answer := models.Answer{
		AnswerUuid: uuid.New(),
		Content: request.Content,
		CreatedAt: time.Now().UnixMilli(),
	}
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: request.QuestionUuid,
		AnswerUuid: answer.AnswerUuid,
	}

	if _, exists := a.QuestionDao.SelectQuestion(request.QuestionUuid); !exists {
		log.Println("Unable to select question")
		return models.CreateAnswerResponse{}, errors.New("unable to select question")
	}

	if err := a.AnswerDao.UpsertAnswer(answer); err != nil {
		log.Println("Unable to upsert answer", err)
		return models.CreateAnswerResponse{}, err
	}

	if err := a.QuestionAnswersDao.InsertQuestionAnswers(questionAnswers); err != nil {
		log.Println("Unable to upsert questionAnswers", err)

		if err := a.AnswerDao.DeleteAnswer(answer.AnswerUuid); err != nil {
			log.Println("Unable to delete upserted answer", err)
		}
		return models.CreateAnswerResponse{}, err
	}

	return models.CreateAnswerResponse{Answer: answer}, nil
}

func (a AnswerManagerImpl) DeleteAnswer(request models.DeleteAnswerRequest) (models.DeleteAnswerResponse, error)  {
	if err := a.QuestionAnswersDao.DeleteQuestionAnswersByAnswerUuid(request.AnswerUuid); err != nil {
		return models.DeleteAnswerResponse{}, errors.New("unable to delete questionAnswers")
	}

	if err := a.AnswerDao.DeleteAnswer(request.AnswerUuid); err != nil {
		return models.DeleteAnswerResponse{}, errors.New("unable to delete answer")
	}

	return models.DeleteAnswerResponse{}, nil
}

func (a AnswerManagerImpl) GetAnswers(request models.GetAnswersRequest) models.GetAnswersResponse {
	questionAnswersS := a.QuestionAnswersDao.SelectQuestionAnswersByQuestionUuid(request.QuestionUuid)

	var answerUuids []uuid.UUID
	for _, questionAnswers := range questionAnswersS {
		answerUuids = append(answerUuids, questionAnswers.AnswerUuid)
	}
	answers := a.AnswerDao.SelectAnswers(answerUuids)

	response := models.GetAnswersResponse{
		Answers: answers,
	}
	return response
}