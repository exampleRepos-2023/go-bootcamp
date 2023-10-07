package managers

import (
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
)

var questionDaoMock = QuestionDaoMock{}
var answerDaoMock = AnswerDaoMock{}
var questionAnswersDaoMock = QuestionAnswersDaoMock{}

type QuestionDaoMock struct {
	Questions   []models.Question
	UpsertError error
	DeleteError error
}

func (q QuestionDaoMock) SelectQuestion(questionUuid uuid.UUID) (models.Question, bool) {
	if len(q.Questions) == 0 {
		return models.Question{}, false
	}
	return q.Questions[0], true
}

func (q QuestionDaoMock) SelectQuestions() []models.Question {
	return q.Questions
}

func (q QuestionDaoMock) UpsertQuestion(question models.Question) error {
	return q.UpsertError
}

func (q QuestionDaoMock) DeleteQuestion(questionUuid uuid.UUID) error {
	return q.DeleteError
}

type AnswerDaoMock struct {
	Answers     []models.Answer
	UpsertError error
	DeleteError error
}

func (a AnswerDaoMock) SelectAnswer(answerUuid uuid.UUID) (models.Answer, bool) {
	for _, answer := range a.Answers {
		if answer.AnswerUuid == answerUuid {
			return answer, true
		}
	}

	return models.Answer{}, false
}

func (a AnswerDaoMock) SelectAnswers(answerUuids []uuid.UUID) []models.Answer {
	answerUuidsMap := map[uuid.UUID]interface{}{}
	for _, answerUuid := range answerUuids {
		answerUuidsMap[answerUuid] = struct{}{}
	}

	var answers []models.Answer
	for _, answer := range a.Answers {
		if answerUuidsMap[answer.AnswerUuid] != nil {
			answers = append(answers, answer)
		}
	}

	return answers
}

func (a AnswerDaoMock) UpsertAnswer(answer models.Answer) error {
	return a.UpsertError
}

func (a AnswerDaoMock) DeleteAnswer(answerUuid uuid.UUID) error {
	return a.DeleteError
}

func (a AnswerDaoMock) DeleteAnswers(answerUuids []uuid.UUID) error {
	return a.DeleteError
}

type QuestionAnswersDaoMock struct {
	QuestionAnswersS []models.QuestionAnswers
	InsertError error
	DeleteError error
}

func (q QuestionAnswersDaoMock) SelectQuestionAnswersByQuestionUuid(questionUuid uuid.UUID) []models.QuestionAnswers {
	return q.QuestionAnswersS
}

func (q QuestionAnswersDaoMock) InsertQuestionAnswers(questionAnswers models.QuestionAnswers) error {
	return q.InsertError
}

func (q QuestionAnswersDaoMock) DeleteQuestionAnswersByQuestionUuid(questionUuid uuid.UUID) error {
	return q.DeleteError
}

func (q QuestionAnswersDaoMock) DeleteQuestionAnswersByAnswerUuid(answerUuid uuid.UUID) error {
	return q.DeleteError
}



