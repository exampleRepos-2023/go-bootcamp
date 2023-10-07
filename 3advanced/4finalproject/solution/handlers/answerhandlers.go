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

var answerHandlers AnswerHandlers = AnswerHandlersImpl{
	AnswerManager: managers.AM,
}

type AnswerHandlers interface {
	handleAnswers(w http.ResponseWriter, r *http.Request)
	handleAnswer(w http.ResponseWriter, r *http.Request)
}

type AnswerHandlersImpl struct {
	AnswerManager managers.AnswerManager
}

func (a AnswerHandlersImpl) handleAnswers(w http.ResponseWriter, r *http.Request) {
	err := validateGetAnswersRequest(r)
	if err != nil {
		log.Println("Unable to validate request", err)
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	request, err := getGetAnswersRequest(r)
	if err != nil {
		log.Println("Unable to get request", err)
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	response := a.AnswerManager.GetAnswers(request)
	writeResponse(w, http.StatusOK, response)
}

func (a AnswerHandlersImpl) handleAnswer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		a.handleCreateAnswer(w, r)
	case "DELETE":
		a.handleDeleteAnswer(w, r)
	default:
		log.Println("Unable to handle answer request")
		writeResponse(w, http.StatusBadRequest, nil)
	}
}

func (a AnswerHandlersImpl) handleCreateAnswer(w http.ResponseWriter, r *http.Request) {
	request, err := getCreateAnswerRequest(r)
	if err != nil {
		log.Println("Unable to get request", err)
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	response, err := a.AnswerManager.CreateAnswer(request)
	if err != nil {
		log.Println("Unable to create answer", err)
		writeResponse(w, http.StatusServiceUnavailable, nil)
		return
	}
	writeResponse(w, http.StatusOK, response)
}

func (a AnswerHandlersImpl) handleDeleteAnswer(w http.ResponseWriter, r *http.Request) {
	request, err := getDeleteAnswerRequest(r)
	if err != nil {
		log.Println("Unable to get answer", err)
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	response, err := a.AnswerManager.DeleteAnswer(request)
	if err != nil {
		log.Println("Unable to delete answer", err)
		writeResponse(w, http.StatusServiceUnavailable, nil)
		return
	}

	writeResponse(w, http.StatusOK, response)
}

func validateGetAnswersRequest(r *http.Request) error {
	if r.Method != "GET" {
		return errors.New("get answers request method isn't POST")
	}
	return nil
}

func getGetAnswersRequest(r *http.Request) (models.GetAnswersRequest, error) {
	var request models.GetAnswersRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read body")
		return models.GetAnswersRequest{}, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Unable to unmarshal request")
		return models.GetAnswersRequest{}, err
	}
	return request, nil
}

func getCreateAnswerRequest(r *http.Request) (models.CreateAnswerRequest, error) {
	var request models.CreateAnswerRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read body")
		return models.CreateAnswerRequest{}, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Unable to unmarshal request")
		return models.CreateAnswerRequest{}, err
	}
	return request, nil
}

func getDeleteAnswerRequest(r *http.Request) (models.DeleteAnswerRequest, error) {
	var request models.DeleteAnswerRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read body")
		return models.DeleteAnswerRequest{}, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Unable to unmarshal request")
		return models.DeleteAnswerRequest{}, err
	}
	return request, nil
}