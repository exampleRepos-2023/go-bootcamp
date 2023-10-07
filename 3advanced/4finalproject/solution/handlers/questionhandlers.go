package handlers

import (
	"encoding/json"
	"errors"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/managers"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"io"
	"log"
	"net/http"
)

var questionHandlers QuestionHandlers = QuestionHandlersImpl{
	QuestionManager: managers.QM,
}

type QuestionHandlers interface {
	handleQuestions(w http.ResponseWriter, r *http.Request)
	handleQuestion(w http.ResponseWriter, r *http.Request)
}

type QuestionHandlersImpl struct {
	QuestionManager managers.QuestionManager
}

func (q QuestionHandlersImpl) handleQuestions(w http.ResponseWriter, r *http.Request) {
	err := validateGetQuestionsRequest(r)
	if err != nil {
		log.Println("Unable to validate request", err)
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}
	response := q.QuestionManager.GetQuestions()
	writeResponse(w, http.StatusOK, response)
}

func (q QuestionHandlersImpl) handleQuestion(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		q.handleCreateQuestion(w, r)
	case "DELETE":
		q.handleDeleteQuestion(w, r)
	default:
		log.Println("Unable to handle question request")
		writeResponse(w, http.StatusBadRequest, nil)
	}
}

func (q QuestionHandlersImpl) handleCreateQuestion(w http.ResponseWriter, r *http.Request) {
	request, err := getCreateQuestionRequest(r)
	if err != nil {
		log.Println("Unable to get request", err)
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	response, err := q.QuestionManager.CreateQuestion(request)
	if err != nil {
		log.Println("Unable to create question", err)
		writeResponse(w, http.StatusServiceUnavailable, nil)
		return
	}
	writeResponse(w, http.StatusOK, response)
}

func (q QuestionHandlersImpl) handleDeleteQuestion(w http.ResponseWriter, r *http.Request) {
	request, err := getDeleteQuestionRequest(r)
	if err != nil {
		log.Println("Unable to get question", err)
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	response, err := q.QuestionManager.DeleteQuestion(request)
	if err != nil {
		log.Println("Unable to delete question", err)
		writeResponse(w, http.StatusServiceUnavailable, nil)
		return
	}

	writeResponse(w, http.StatusOK, response)
}

func validateGetQuestionsRequest(r *http.Request) error {
	if r.Method != "GET" {
		return errors.New("get questions request method isn't POST")
	}
	return nil
}

func getCreateQuestionRequest(r *http.Request) (models.CreateQuestionRequest, error) {
	var request models.CreateQuestionRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read body")
		return models.CreateQuestionRequest{}, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Unable to unmarshal quest")
		return models.CreateQuestionRequest{}, err
	}
	return request, nil
}

func getDeleteQuestionRequest(r *http.Request) (models.DeleteQuestionRequest, error) {
	var request models.DeleteQuestionRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read body")
		return models.DeleteQuestionRequest{}, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Unable to unmarshal request")
		return models.DeleteQuestionRequest{}, err
	}
	return request, nil
}