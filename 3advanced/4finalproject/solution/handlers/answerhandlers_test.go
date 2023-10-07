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

func TestHandleAnswers_GetAnswers_Success(t *testing.T) {
	question := models.Question{
		QuestionUuid: uuid.New(),
		Title: "Title",
		Description: "Description",
	}
	answer := models.Answer{
		AnswerUuid: uuid.New(),
		Content: "Content",
	}
	answers := []models.Answer{
		answer,
	}
	getAnswersRequest := models.GetAnswersRequest{
		QuestionUuid: question.QuestionUuid,
	}
	answerManagerMock.GetAnswersRequest = getAnswersRequest
	answerManagerMock.GetAnswersResponse = models.GetAnswersResponse{
		Answers: answers,
	}
	answerManagerMock.GetAnswersCalled = false

	answerHandlers = AnswerHandlersImpl{
		AnswerManager: &answerManagerMock,
	}

	requestJson, err := json.Marshal(getAnswersRequest)
	if err != nil {
		t.Fatal("Unable to marshal request", err)
	}
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(
		"GET",
		"http://localhost/answers",
		strings.NewReader(string(requestJson)))
	request.Header.Set("Accept", "application/json")

	answerHandlers.handleAnswers(writer, request)

	response := writer.Result()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Unable read response body", err)
	}

	var unmarshalledResponse models.GetAnswersResponse
	err = json.Unmarshal(responseBody, &unmarshalledResponse)
	if err != nil {
		t.Fatal("Unable to unmarshal response")
		return
	}

	if len(unmarshalledResponse.Answers) != 1 {
		t.Fatal("Unable to get correct number of answers")
	}

	if unmarshalledResponse.Answers[0] != answer {
		t.Error("Unable to get correct answers")
	}

	if !answerManagerMock.GetAnswersCalled {
		t.Error("Unable to call answerManager")
	}
}

func TestHandleAnswers_GetAnswers_NotFound(t *testing.T) {
	question := models.Question{
		QuestionUuid: uuid.New(),
		Title: "Title",
		Description: "Description",
	}
	var answers []models.Answer
	getAnswersRequest := models.GetAnswersRequest{
		QuestionUuid: question.QuestionUuid,
	}
	answerManagerMock.GetAnswersRequest = getAnswersRequest
	answerManagerMock.GetAnswersResponse = models.GetAnswersResponse{
		Answers: answers,
	}
	answerManagerMock.GetAnswersCalled = false

	answerHandlers = AnswerHandlersImpl{
		AnswerManager: &answerManagerMock,
	}

	requestJson, err := json.Marshal(getAnswersRequest)
	if err != nil {
		t.Fatal("Unable to marshal request", err)
	}
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(
		"GET",
		"http://localhost/answers",
		strings.NewReader(string(requestJson)))
	request.Header.Set("Accept", "application/json")

	answerHandlers.handleAnswers(writer, request)

	response := writer.Result()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Unable read response body", err)
	}

	var unmarshalledResponse models.GetAnswersResponse
	err = json.Unmarshal(responseBody, &unmarshalledResponse)
	if err != nil {
		t.Fatal("Unable to unmarshal response")
		return
	}

	if len(unmarshalledResponse.Answers) != 0 {
		t.Fatal("Unable to get correct number of answers")
	}

	if !answerManagerMock.GetAnswersCalled {
		t.Error("Unable to call answerManager")
	}
}

func TestHandleAnswer_CreateAnswer_Success(t *testing.T) {
	questionUuid := uuid.New()
	createdAnswer := models.Answer{
		AnswerUuid: uuid.New(),
		Content: "Content",
	}
	createAnswerRequest := models.CreateAnswerRequest{
		QuestionUuid: questionUuid,
		Content: "Content",
	}
	answerManagerMock.CreateAnswerRequest = createAnswerRequest
	answerManagerMock.CreateAnswerResponse = models.CreateAnswerResponse{
		Answer: createdAnswer,
	}
	answerManagerMock.CreateAnswerError = nil
	answerManagerMock.CreateAnswerCalled = false

	answerHandlers = AnswerHandlersImpl{
		AnswerManager: &answerManagerMock,
	}

	writer := httptest.NewRecorder()
	createRequestJson, err := json.Marshal(createAnswerRequest)
	if err != nil {
		t.Fatal("Unable to marshal answer", err)
	}
	request := httptest.NewRequest(
		"POST",
		"http://localhost/answer",
		strings.NewReader(string(createRequestJson)))
	request.Header.Set("Accept", "application/json")

	answerHandlers.handleAnswer(writer, request)

	response := writer.Result()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Unable read response body", err)
	}

	var unmarshalledResponse models.CreateAnswerResponse
	err = json.Unmarshal(responseBody, &unmarshalledResponse)
	if err != nil {
		t.Fatal("Unable to unmarshal response")
		return
	}
	if unmarshalledResponse.Answer != createdAnswer {
		t.Error("Unable to get correct answers")
	}

	if !answerManagerMock.CreateAnswerCalled {
		t.Error("Unable to call answerManager")
	}
}

func TestHandleAnswer_CreateAnswer_Error(t *testing.T) {
	questionUuid := uuid.New()
	createdAnswer := models.Answer{
		AnswerUuid: uuid.New(),
		Content: "Content",
	}
	createAnswerRequest := models.CreateAnswerRequest{
		QuestionUuid: questionUuid,
		Content: "Content",
	}
	answerManagerMock.CreateAnswerRequest = createAnswerRequest
	answerManagerMock.CreateAnswerResponse = models.CreateAnswerResponse{
		Answer: createdAnswer,
	}
	answerManagerMock.CreateAnswerError = errors.New("unable to create answer")
	answerManagerMock.CreateAnswerCalled = false

	answerHandlers = AnswerHandlersImpl{
		AnswerManager: &answerManagerMock,
	}

	writer := httptest.NewRecorder()
	createRequestJson, err := json.Marshal(createAnswerRequest)
	if err != nil {
		t.Fatal("Unable to marshal answer", err)
	}
	request := httptest.NewRequest(
		"POST",
		"http://localhost/answer",
		strings.NewReader(string(createRequestJson)))
	request.Header.Set("Accept", "application/json")

	answerHandlers.handleAnswer(writer, request)

	response := writer.Result()
	if response.StatusCode != http.StatusServiceUnavailable {
		t.Error("Unable to detect create answer error")
	}

	if !answerManagerMock.CreateAnswerCalled {
		t.Error("Unable to call answerManager")
	}
}

func TestHandleAnswer_DeleteAnswer_Success(t *testing.T) {
	deleteAnswerRequest := models.DeleteAnswerRequest{
		AnswerUuid: uuid.New(),
	}
	answerManagerMock.DeleteAnswerRequest = deleteAnswerRequest
	answerManagerMock.DeleteAnswerError = nil
	answerManagerMock.DeleteAnswerCalled = false

	answerHandlers = AnswerHandlersImpl{
		AnswerManager: &answerManagerMock,
	}

	writer := httptest.NewRecorder()
	deleteRequestJson, err := json.Marshal(deleteAnswerRequest)
	if err != nil {
		t.Fatal("Unable to marshal answer", err)
	}
	request := httptest.NewRequest(
		"DELETE",
		"http://localhost/answer",
		strings.NewReader(string(deleteRequestJson)))
	request.Header.Set("Accept", "application/json")

	answerHandlers.handleAnswer(writer, request)

	response := writer.Result()

	if response.StatusCode != http.StatusOK {
		t.Error("Unable to delete answer")
	}
}

func TestHandleAnswer_DeleteAnswer_Error(t *testing.T) {
	deleteAnswerRequest := models.DeleteAnswerRequest{
		AnswerUuid: uuid.New(),
	}
	answerManagerMock.DeleteAnswerRequest = deleteAnswerRequest
	answerManagerMock.DeleteAnswerError = errors.New("unable to delete answer")
	answerManagerMock.DeleteAnswerCalled = false

	answerHandlers = AnswerHandlersImpl{
		AnswerManager: &answerManagerMock,
	}

	writer := httptest.NewRecorder()
	deleteRequestJson, err := json.Marshal(deleteAnswerRequest)
	if err != nil {
		t.Fatal("Unable to marshal answer", err)
	}
	request := httptest.NewRequest(
		"POST",
		"http://localhost/answer",
		strings.NewReader(string(deleteRequestJson)))
	request.Header.Set("Accept", "application/json")

	answerHandlers.handleAnswer(writer, request)

	response := writer.Result()

	if response.StatusCode != http.StatusServiceUnavailable {
		t.Error("Unable to detect delete answer error")
	}
}