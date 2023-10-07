package managers

import (
	"errors"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"testing"
)

func TestAnswerManagerImpl_CreateAnswer_Success(t *testing.T) {
	question := models.Question{
		QuestionUuid: uuid.New(),
	}
	questionDaoMock.Questions = []models.Question{question}
	answerDaoMock.UpsertError = nil
	questionAnswersDaoMock.InsertError = nil
	request := models.CreateAnswerRequest{
		QuestionUuid: question.QuestionUuid,
		Content: "Content",
	}

	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}

	response, err := answerManager.CreateAnswer(request)
	if err != nil {
		t.Error("Unable to create answer")
	}
	if response.Answer.Content != request.Content {
		t.Error("Unable to create answer title")
	}
}

func TestAnswerManagerImpl_CreateAnswer_SelectQuestionNotFound(t *testing.T) {
	questionDaoMock.Questions = []models.Question{}
	answerDaoMock.UpsertError = nil
	questionAnswersDaoMock.InsertError = nil
	request := models.CreateAnswerRequest{
		QuestionUuid: uuid.New(),
		Content: "Content",
	}

	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}

	_, err := answerManager.CreateAnswer(request)
	if err == nil {
		t.Error("Unable to detect empty question")
	}
}

func TestAnswerManagerImpl_CreateAnswer_InsertAnswerError(t *testing.T) {
	question := models.Question{
		QuestionUuid: uuid.New(),
	}
	questionDaoMock.Questions = []models.Question{question}
	createAnswerError := "unable to create answer"
	answerDaoMock.UpsertError = errors.New(createAnswerError)
	questionAnswersDaoMock.InsertError = nil
	request := models.CreateAnswerRequest{
		QuestionUuid: question.QuestionUuid,
		Content: "Content",
	}

	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}

	_, err := answerManager.CreateAnswer(request)
	if err == nil {
		t.Error("Unable to detect insert answer error")
	}
}

func TestAnswerManagerImpl_CreateAnswer_InsertQuestionAnswersError(t *testing.T) {
	question := models.Question{
		QuestionUuid: uuid.New(),
	}
	questionDaoMock.Questions = []models.Question{question}
	answerDaoMock.UpsertError = nil
	questionAnswersDaoMock.InsertError = errors.New("unable to create questionAnswers")
	request := models.CreateAnswerRequest{
		QuestionUuid: question.QuestionUuid,
		Content: "Content",
	}

	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}

	_, err := answerManager.CreateAnswer(request)
	if err == nil {
		t.Error("Unable to detect insert questionAnswers error")
	}
}

func TestAnswerManagerImpl_GetAnswers_Success(t *testing.T) {
	questionUuid := uuid.New()
	answerUuid1 := uuid.New()
	answerUuid2 := uuid.New()
	answerDaoMock.Answers = []models.Answer{
		{AnswerUuid: answerUuid1},
		{AnswerUuid: answerUuid2},
	}
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{
		{QuestionUuid: questionUuid, AnswerUuid: answerUuid1},
		{QuestionUuid: questionUuid, AnswerUuid: answerUuid2},
	}
	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.GetAnswersRequest{
		QuestionUuid: questionUuid,
	}

	response := answerManager.GetAnswers(request)
	if len(response.Answers) != 2 {
		t.Error("Unable to get answers correctly")
	}
	if response.Answers[0].AnswerUuid != answerUuid1 && response.Answers[1].AnswerUuid != answerUuid1 {
		t.Error("Unable to get answer1")
	}
	if response.Answers[0].AnswerUuid != answerUuid2 && response.Answers[1].AnswerUuid != answerUuid2 {
		t.Error("Unable to get answer2")
	}
}

func TestAnswerManagerImpl_GetAnswers_EmptyQuestionAnswers(t *testing.T) {
	questionUuid := uuid.New()
	answerUuid1 := uuid.New()
	answerUuid2 := uuid.New()
	answerDaoMock.Answers = []models.Answer{
		{AnswerUuid: answerUuid1},
		{AnswerUuid: answerUuid2},
	}
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{}
	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.GetAnswersRequest{
		QuestionUuid: questionUuid,
	}

	response := answerManager.GetAnswers(request)
	if len(response.Answers) != 0 {
		t.Error("Unable to get empty answers")
	}
}

func TestAnswerManagerImpl_GetAnswers_EmptyAnswer(t *testing.T) {
	questionUuid := uuid.New()
	answerUuid1 := uuid.New()
	answerUuid2 := uuid.New()
	answerDaoMock.Answers = []models.Answer{}
	questionAnswersDaoMock.QuestionAnswersS = []models.QuestionAnswers{
		{QuestionUuid: questionUuid, AnswerUuid: answerUuid1},
		{QuestionUuid: questionUuid, AnswerUuid: answerUuid2},
	}
	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.GetAnswersRequest{
		QuestionUuid: questionUuid,
	}

	response := answerManager.GetAnswers(request)
	if len(response.Answers) != 0 {
		t.Error("Unable to get empty answers")
	}
}

func TestAnswerManagerImpl_DeleteAnswer_Success(t *testing.T) {
	answerDaoMock.DeleteError = nil
	questionAnswersDaoMock.DeleteError = nil
	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.DeleteAnswerRequest{
		AnswerUuid: uuid.New(),
	}
	_, err := answerManager.DeleteAnswer(request)
	if err != nil {
		t.Error("Unable to delete answer")
	}
}

func TestAnswerManagerImpl_DeleteAnswer_DeleteQuestionAnswersError(t *testing.T) {
	answerDaoMock.DeleteError = nil
	questionAnswersDaoMock.DeleteError = errors.New("unable to delete questionAnswers")
	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.DeleteAnswerRequest{
		AnswerUuid: uuid.New(),
	}
	_, err := answerManager.DeleteAnswer(request)
	if err == nil {
		t.Error("Unable to detect delete questionAnswers error")
	}
}

func TestAnswerManagerImpl_DeleteAnswer_DeleteAnswerError(t *testing.T) {
	answerDaoMock.DeleteError = errors.New("unable to delete answer")
	questionAnswersDaoMock.DeleteError = nil
	answerManager := AnswerManagerImpl{
		QuestionDao: questionDaoMock,
		AnswerDao: answerDaoMock,
		QuestionAnswersDao: questionAnswersDaoMock,
	}
	request := models.DeleteAnswerRequest{
		AnswerUuid: uuid.New(),
	}
	_, err := answerManager.DeleteAnswer(request)
	if err == nil {
		t.Error("Unable to detect delete questionAnswers error")
	}
}
