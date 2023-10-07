package handlers

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleQuestions_GetQuestions_Success(t *testing.T) {
	question := models.Question{
		QuestionUuid: uuid.New(),
		Title: "Title",
		Description: "Description",
	}
	questions := []models.Question{
		question,
	}
	questionManagerMock.GetQuestionsResponse = models.GetQuestionsResponse{
		Questions: questions,
	}
	questionManagerMock.GetQuestionsCalled = false

	questionHandlers = QuestionHandlersImpl{
		QuestionManager: &questionManagerMock,
	}

	writer := httptest.NewRecorder()
	request := httptest.NewRequest(
		"GET",
		"http://localhost/questions",
		strings.NewReader(""))
	request.Header.Set("Accept", "application/json")

	questionHandlers.handleQuestions(writer, request)

	response := writer.Result()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Unable read response body", err)
	}

	var unmarshalledResponse models.GetQuestionsResponse
	err = json.Unmarshal(responseBody, &unmarshalledResponse)
	if err != nil {
		t.Fatal("Unable to unmarshal response")
		return
	}

	if len(unmarshalledResponse.Questions) != 1 {
		t.Fatal("Unable to get correct number of questions")
	}

	if unmarshalledResponse.Questions[0] != question {
		t.Error("Unable to get correct questions")
	}

	if !questionManagerMock.GetQuestionsCalled {
		t.Error("Unable to call questionManager")
	}
}

func TestHandleQuestions_GetQuestions_NotFound(t *testing.T) {
	var questions []models.Question
	questionManagerMock.GetQuestionsResponse = models.GetQuestionsResponse{
		Questions: questions,
	}
	questionManagerMock.GetQuestionsCalled = false

	questionHandlers = QuestionHandlersImpl{
		QuestionManager: &questionManagerMock,
	}

	writer := httptest.NewRecorder()
	request := httptest.NewRequest(
		"GET",
		"http://localhost/questions",
		strings.NewReader(""))
	request.Header.Set("Accept", "application/json")

	questionHandlers.handleQuestions(writer, request)

	response := writer.Result()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Unable read response body", err)
	}

	var unmarshalledResponse models.GetQuestionsResponse
	err = json.Unmarshal(responseBody, &unmarshalledResponse)
	if err != nil {
		t.Fatal("Unable to unmarshal response")
		return
	}

	if len(unmarshalledResponse.Questions) != 0 {
		t.Fatal("Unable to get correct number of questions")
	}

	if !questionManagerMock.GetQuestionsCalled {
		t.Error("Unable to call questionManager")
	}
}

func TestHandleQuestion_CreateQuestion_Success(t *testing.T) {
	createdQuestion := models.Question{
		QuestionUuid: uuid.New(),
		Title: "Title",
		Description: "Description",
	}
	createQuestionRequest := models.CreateQuestionRequest{
		Title: createdQuestion.Title,
		Description: createdQuestion.Description,
	}
	questionManagerMock.CreateQuestionRequest = createQuestionRequest
	questionManagerMock.CreateQuestionResponse = models.CreateQuestionResponse{
		Question: createdQuestion,
	}
	questionManagerMock.CreateQuestionError = nil
	questionManagerMock.CreateQuestionCalled = false

	questionHandlers = QuestionHandlersImpl{
		QuestionManager: &questionManagerMock,
	}

	writer := httptest.NewRecorder()
	createRequestJson, err := json.Marshal(createQuestionRequest)
	if err != nil {
		t.Fatal("Unable to marshal question", err)
	}
	request := httptest.NewRequest(
		"POST",
		"http://localhost/question",
		strings.NewReader(string(createRequestJson)))
	request.Header.Set("Accept", "application/json")

	questionHandlers.handleQuestion(writer, request)

	response := writer.Result()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Unable read response body", err)
	}

	var unmarshalledResponse models.CreateQuestionResponse
	err = json.Unmarshal(responseBody, &unmarshalledResponse)
	if err != nil {
		t.Fatal("Unable to unmarshal response")
		return
	}
	if unmarshalledResponse.Question != createdQuestion {
		t.Error("Unable to get correct questions")
	}

	if !questionManagerMock.CreateQuestionCalled {
		t.Error("Unable to call questionManager")
	}
}

func TestHandleQuestion_CreateQuestion_Error(t *testing.T) {
	createdQuestion := models.Question{
		QuestionUuid: uuid.New(),
		Title: "Title",
		Description: "Description",
	}
	createQuestionRequest := models.CreateQuestionRequest{
		Title: createdQuestion.Title,
		Description: createdQuestion.Description,
	}
	questionManagerMock.CreateQuestionRequest = createQuestionRequest
	questionManagerMock.CreateQuestionResponse = models.CreateQuestionResponse{
		Question: createdQuestion,
	}
	questionManagerMock.CreateQuestionError = errors.New("unable to create question")
	questionManagerMock.CreateQuestionCalled = false

	questionHandlers = QuestionHandlersImpl{
		QuestionManager: &questionManagerMock,
	}

	writer := httptest.NewRecorder()
	createRequestJson, err := json.Marshal(createQuestionRequest)
	if err != nil {
		t.Fatal("Unable to marshal question", err)
	}
	request := httptest.NewRequest(
		"POST",
		"http://localhost/question",
		strings.NewReader(string(createRequestJson)))
	request.Header.Set("Accept", "application/json")

	questionHandlers.handleQuestion(writer, request)

	response := writer.Result()
	if response.StatusCode != http.StatusServiceUnavailable {
		t.Error("Unable to detect create question error")
	}

	if !questionManagerMock.CreateQuestionCalled {
		t.Error("Unable to call questionManager")
	}
}

func TestHandleQuestion_DeleteQuestion_Success(t *testing.T) {
	deleteQuestionRequest := models.DeleteQuestionRequest{
		QuestionUuid: uuid.New(),
	}
	questionManagerMock.DeleteQuestionRequest = deleteQuestionRequest
	questionManagerMock.DeleteQuestionError = nil
	questionManagerMock.DeleteQuestionCalled = false

	questionHandlers = QuestionHandlersImpl{
		QuestionManager: &questionManagerMock,
	}

	writer := httptest.NewRecorder()
	deleteRequestJson, err := json.Marshal(deleteQuestionRequest)
	if err != nil {
		t.Fatal("Unable to marshal question", err)
	}
	request := httptest.NewRequest(
		"DELETE",
		"http://localhost/question",
		strings.NewReader(string(deleteRequestJson)))
	request.Header.Set("Accept", "application/json")

	questionHandlers.handleQuestion(writer, request)

	response := writer.Result()

	if response.StatusCode != http.StatusOK {
		t.Error("Unable to delete question")
	}
}

func TestHandleQuestion_DeleteQuestion_Error(t *testing.T) {
	deleteQuestionRequest := models.DeleteQuestionRequest{
		QuestionUuid: uuid.New(),
	}
	questionManagerMock.DeleteQuestionRequest = deleteQuestionRequest
	questionManagerMock.DeleteQuestionError = errors.New("unable to delete question")
	questionManagerMock.DeleteQuestionCalled = false

	questionHandlers = QuestionHandlersImpl{
		QuestionManager: &questionManagerMock,
	}

	writer := httptest.NewRecorder()
	deleteRequestJson, err := json.Marshal(deleteQuestionRequest)
	if err != nil {
		t.Fatal("Unable to marshal question", err)
	}
	request := httptest.NewRequest(
		"POST",
		"http://localhost/question",
		strings.NewReader(string(deleteRequestJson)))
	request.Header.Set("Accept", "application/json")

	questionHandlers.handleQuestion(writer, request)

	response := writer.Result()

	if response.StatusCode != http.StatusServiceUnavailable {
		t.Error("Unable to detect delete question error")
	}
}