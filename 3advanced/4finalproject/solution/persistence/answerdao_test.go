package persistence

import (
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"testing"
	"time"
)

func TestAnswerDaoImpl_SelectAnswers(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	answerUuid := uuid.New()
	content := "Content"
	now := time.Now().UnixMilli()
	answer := models.Answer{
		AnswerUuid: answerUuid,
		Content: content,
		CreatedAt: now,
	}
	diffAnswerUuid := uuid.New()
	diffAnswer := models.Answer{
		AnswerUuid: diffAnswerUuid,
		Content: content,
		CreatedAt: now,
	}
	answerUuids := []uuid.UUID{answer.AnswerUuid, diffAnswer.AnswerUuid}

	answers := Answers.SelectAnswers(answerUuids)
	if len(answers) != 0 {
		t.Error("Unable to start with empty database")
	}

	err := Answers.UpsertAnswer(answer)
	if err != nil {
		t.Error("Unable to upsert answer")
	}

	err = Answers.UpsertAnswer(diffAnswer)
	if err != nil {
		t.Error("Unable to upsert a different answer")
	}

	answers = Answers.SelectAnswers(answerUuids)
	if len(answers) != 2 {
		t.Error("Unable to insert two answers")
	}
	if answers[0] != answer && answers[0] != diffAnswer {
		t.Error("Unable to insert one answer correctly")
	}
	if answers[1] != answer && answers[1] != diffAnswer {
		t.Error("Unable to insert a different answer correctly")
	}
}

func TestAnswerDaoImpl_SelectAnswer(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	answerUuid := uuid.New()
	content := "Content"
	now := time.Now().UnixMilli()
	answer := models.Answer{
		AnswerUuid: answerUuid,
		Content: content,
		CreatedAt: now,
	}

	_, exists := Answers.SelectAnswer(answer.AnswerUuid)
	if exists {
		t.Error("Unable to start with empty database")
	}

	err := Answers.UpsertAnswer(answer)
	if err != nil {
		t.Error("Unable to insert answer")
	}

	insertedAnswer, exists := Answers.SelectAnswer(answer.AnswerUuid)
	if !exists {
		t.Error("Unable to select answer")
	}
	if insertedAnswer != answer {
		t.Error("Unable to select correct answer")
	}
}

func TestAnswerDaoImpl_UpsertAnswer(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	answerUuid := uuid.New()
	content := "Content"
	now := time.Now().UnixMilli()
	answer := models.Answer{
		AnswerUuid: answerUuid,
		Content: content,
		CreatedAt: now,
	}
	diffAnswerUuid := uuid.New()
	diffAnswer := models.Answer{
		AnswerUuid: diffAnswerUuid,
		Content: content,
		CreatedAt: now,
	}
	answerUuids := []uuid.UUID{answer.AnswerUuid, diffAnswer.AnswerUuid}

	answers := Answers.SelectAnswers(answerUuids)
	if len(answers) != 0 {
		t.Error("Unable to start with empty database")
	}

	err := Answers.UpsertAnswer(answer)
	if err != nil {
		t.Error("Unable to insert answer")
	}

	insertedAnswer, exists := Answers.SelectAnswer(answer.AnswerUuid)
	if !exists {
		t.Error("Unable to insert answer")
	}
	if insertedAnswer != answer {
		t.Error("Unable to insert answer correctly")
	}

	updatedContent := "Updated content"
	updatedNow := time.UnixMilli(now).Add(time.Second).UnixMilli()
	updatedAnswer := models.Answer{
		AnswerUuid: answerUuid,
		Content: updatedContent,
		CreatedAt: updatedNow,
	}
	err = Answers.UpsertAnswer(updatedAnswer)
	if err != nil {
		t.Error("Unable to update answer")
	}

	updatedAnswer, exists = Answers.SelectAnswer(answerUuid)
	if !exists {
		t.Error("Unable to keep inserted answer when updating")
	}
	if updatedAnswer.AnswerUuid != answer.AnswerUuid {
		t.Error("Unable to keep answer uuid")
	}
	if updatedAnswer.Content == answer.Content {
		t.Error("Unable to update answer content")
	}
	if updatedAnswer.CreatedAt == answer.CreatedAt {
		t.Error("Unable to update answer createdAt")
	}

	err = Answers.UpsertAnswer(diffAnswer)
	if err != nil {
		t.Error("Unable to insert a different answer")
	}

	diffInsertedAnswer, exists := Answers.SelectAnswer(diffAnswerUuid)
	if !exists {
		t.Error("Unable to insert a different answer")
	}
	if diffInsertedAnswer != diffAnswer {
		t.Error("Unable to insert a different answer correctly")
	}

	bothAnswers := Answers.SelectAnswers(answerUuids)
	if len(bothAnswers) != 2 {
		t.Error("Unable to keep two answers")
	}
}

func TestAnswerDaoImpl_DeleteAnswer(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	answerUuid := uuid.New()
	content := "Content"
	now := time.Now().UnixMilli()
	answer := models.Answer{
		AnswerUuid: answerUuid,
		Content: content,
		CreatedAt: now,
	}
	answerUuids := []uuid.UUID{answer.AnswerUuid}

	answers := Answers.SelectAnswers(answerUuids)
	if len(answers) != 0 {
		t.Error("Unable to start with empty database")
	}

	err := Answers.UpsertAnswer(answer)
	if err != nil {
		t.Error("Unable to insert answer")
	}

	answers = Answers.SelectAnswers(answerUuids)
	if len(answers) != 1 {
		t.Error("Unable to insert correctly")
	}

	err = Answers.DeleteAnswer(answer.AnswerUuid)
	if err != nil {
		t.Error("Unable to delete answer")
	}

	answers = Answers.SelectAnswers(answerUuids)
	if len(answers) != 0 {
		t.Error("Unable to delete answer correctly")
	}
}