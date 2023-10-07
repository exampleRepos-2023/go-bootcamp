package managers

import (
	"errors"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"testing"
)

func TestQuestionManagerImpl_CreateQuestion(t *testing.T) {
	questionDaoMock.UpsertError = nil
	answerDaoMock.UpsertError = nil
	questionAnswersDaoMock.InsertError = nil
	questionManager := QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.CreateQuestionRequest{
		Title: "Title",
		Description: "Description",
	}

	response, err := questionManager.CreateQuestion(request)
	if err != nil {
		t.Error("Unable to create question")
	}
	if response.Question.Title != request.Title {
		t.Error("Unable to create question title")
	}
	if response.Question.Description != request.Description {
		t.Error("Unable to create question description")
	}

	createQuestionError := "unable to insert question"
	questionDaoMock.UpsertError = errors.New(createQuestionError)
	questionManager = QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	response, err = questionManager.CreateQuestion(request)
	if err == nil {
		t.Error("Unable to get insert question error")
	}
	if err.Error() != createQuestionError {
		t.Error("Unable to match insert question error")
	}
}

func TestQuestionManagerImpl_GetQuestions(t *testing.T) {
	questionDaoMock.Questions = []models.Question{}
	questionManager := QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	response := questionManager.GetQuestions()
	if len(response.Questions) != 0 {
		t.Error("Unable to get empty questions")
	}

	question := models.Question{
		QuestionUuid: uuid.New(),
	}
	questionDaoMock.Questions = []models.Question{question}
	questionManager = QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	response = questionManager.GetQuestions()
	if len(response.Questions) != 1 {
		t.Error("Unable to one question")
	}
	if response.Questions[0] != question {
		t.Error("Unable to match question")
	}
}

func TestQuestionManagerImpl_DeleteQuestion_Success(t *testing.T) {
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: uuid.New(),
		AnswerUuid: uuid.New(),
	}
	questionDaoMock.DeleteError = nil
	answerDaoMock.DeleteError = nil
	questionAnswersDaoMock.DeleteError = nil
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{questionAnswers}
	questionManager := QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.DeleteQuestionRequest{
		QuestionUuid: questionAnswers.QuestionUuid,
	}
	_, err := questionManager.DeleteQuestion(request)
	if err != nil {
		t.Error("Unable to create question")
	}
}

func TestQuestionManagerImpl_DeleteQuestion_QuestionAnswersNotFound(t *testing.T) {
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: uuid.New(),
		AnswerUuid: uuid.New(),
	}
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{}
	questionDaoMock.DeleteError = nil
	answerDaoMock.DeleteError = nil
	questionAnswersDaoMock.DeleteError = nil
	request := models.DeleteQuestionRequest{
		QuestionUuid: questionAnswers.QuestionUuid,
	}

	questionManager := QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}

	_, err := questionManager.DeleteQuestion(request)
	if err != nil {
		t.Error("Unable to bypass with empty questionAnswers")
	}
}

func TestQuestionManagerImpl_DeleteQuestion_DeleteAnswerError(t *testing.T) {
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: uuid.New(),
		AnswerUuid: uuid.New(),
	}
	deleteAnswerError := errors.New("unable to delete answer")
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{questionAnswers}
	questionDaoMock.DeleteError = nil
	answerDaoMock.DeleteError = deleteAnswerError
	questionAnswersDaoMock.DeleteError = nil
	request := models.DeleteQuestionRequest{
		QuestionUuid: questionAnswers.QuestionUuid,
	}

	questionManager := QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}

	_, err := questionManager.DeleteQuestion(request)
	if err == nil {
		t.Error("Unable to detect delete answer error")
	}
}

func TestQuestionManagerImpl_DeleteQuestion_DeleteQuestionAnswersError(t *testing.T) {
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: uuid.New(),
		AnswerUuid: uuid.New(),
	}
	deleteQuestionAnswersError := errors.New("unable to delete questionAnswers")
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{questionAnswers}
	questionDaoMock.DeleteError = nil
	answerDaoMock.DeleteError = nil
	request := models.DeleteQuestionRequest{
		QuestionUuid: questionAnswers.QuestionUuid,
	}

	questionAnswersDaoMock.DeleteError = deleteQuestionAnswersError
	questionManager := QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	_, err := questionManager.DeleteQuestion(request)
	if err == nil {
		t.Error("Unable to detect delete questionAnswers error")
	}
}

func TestQuestionManagerImpl_DeleteQuestion_DeleteQuestionError(t *testing.T) {
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: uuid.New(),
		AnswerUuid: uuid.New(),
	}
	deleteQuestionError := errors.New("unable to delete question")
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{questionAnswers}
	questionDaoMock.DeleteError = deleteQuestionError
	answerDaoMock.DeleteError = nil
	questionAnswersDaoMock.DeleteError = nil
	request := models.DeleteQuestionRequest{
		QuestionUuid: questionAnswers.QuestionUuid,
	}

	questionManager := QuestionManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	_, err := questionManager.DeleteQuestion(request)
	if err == nil {
		t.Error("Unable to detect delete question error")
	}
}