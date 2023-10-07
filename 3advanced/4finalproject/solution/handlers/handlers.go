package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	var handler http.ServeMux
	handler.HandleFunc("/questions", questionHandlers.handleQuestions)
	handler.HandleFunc("/question", questionHandlers.handleQuestion)
	handler.HandleFunc("/answers", answerHandlers.handleAnswers)
	handler.HandleFunc("/answer", answerHandlers.handleAnswer)

	server := http.Server{
		Addr:         "", // localhost:80 (host:port)
		Handler:      &handler,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Unable to listen and serve")
		return
	}
}

func writeResponse(w http.ResponseWriter, status int, response interface{}) {
	if response == nil {
		w.WriteHeader(status)
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println("Unable to marshal response", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	_, err = fmt.Fprint(w, string(responseJson))
	if err != nil {
		log.Println("Unable to write response", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
