package persistence

import (
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"testing"
	"time"
)

func TestQuestionDaoImpl_SelectQuestions(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	questions := Questions.SelectQuestions()
	if len(questions) != 0 {
		t.Error("Unable to start with empty database")
	}

	questionUuid := uuid.New()
	title := "Title"
	description := "Description"
	now := time.Now().UnixMilli()
	question := models.Question{
		QuestionUuid: questionUuid,
		Title: title,
		Description: description,
		CreatedAt: now,
	}
	err := Questions.UpsertQuestion(question)
	if err != nil {
		t.Error("Unable to upsert question")
	}

	questions = Questions.SelectQuestions()
	if len(questions) != 1 {
		t.Error("Unable to insert one question")
	}
	if questions[0] != question {
		t.Error("Unable to insert question correctly")
	}
}

func TestQuestionDaoImpl_SelectQuestion(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	questionUuid := uuid.New()
	title := "Title"
	description := "Description"
	now := time.Now().UnixMilli()
	question := models.Question{
		QuestionUuid: questionUuid,
		Title: title,
		Description: description,
		CreatedAt: now,
	}

	_, exists := Questions.SelectQuestion(question.QuestionUuid)
	if exists {
		t.Error("Unable to start with empty database")
	}

	err := Questions.UpsertQuestion(question)
	if err != nil {
		t.Error("Unable to insert question")
	}

	insertedQuestion, exists := Questions.SelectQuestion(question.QuestionUuid)
	if !exists {
		t.Error("Unable to select question")
	}
	if insertedQuestion != question {
		t.Error("Unable to select correct question")
	}
}

func TestQuestionDaoImpl_UpsertQuestion(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	questions := Questions.SelectQuestions()
	if len(questions) != 0 {
		t.Error("Unable to start with empty database")
	}

	questionUuid := uuid.New()
	title := "Title"
	description := "Description"
	now := time.Now().UnixMilli()
	question := models.Question{
		QuestionUuid: questionUuid,
		Title: title,
		Description: description,
		CreatedAt: now,
	}

	_, exists := Questions.SelectQuestion(questionUuid)
	if exists {
		t.Error("Unable to start with empty database")
	}
	err := Questions.UpsertQuestion(question)
	if err != nil {
		t.Error("Unable to insert question")
	}

	insertedQuestion, exists := Questions.SelectQuestion(question.QuestionUuid)
	if !exists {
		t.Error("Unable to insert question")
	}
	if insertedQuestion != question {
		t.Error("Unable to insert question correctly")
	}

	updatedTitle := "Updated title"
	updatedDescription := "Updated description"
	updatedNow := time.UnixMilli(now).Add(time.Second).UnixMilli()
	updatedQuestion := models.Question{
		QuestionUuid: questionUuid,
		Title: updatedTitle,
		Description: updatedDescription,
		CreatedAt: updatedNow,
	}
	err = Questions.UpsertQuestion(updatedQuestion)
	if err != nil {
		t.Error("Unable to update question")
	}

	updatedQuestion, exists = Questions.SelectQuestion(questionUuid)
	if !exists {
		t.Error("Unable to keep inserted question when updating")
	}
	if updatedQuestion.QuestionUuid != question.QuestionUuid {
		t.Error("Unable to keep question uuid")
	}
	if updatedQuestion.Title == question.Title {
		t.Error("Unable to update question title")
	}
	if updatedQuestion.Description == question.Description {
		t.Error("Unable to update question description")
	}
	if updatedQuestion.CreatedAt == question.CreatedAt {
		t.Error("Unable to update question createdAt")
	}

	diffQuestionUuid := uuid.New()
	diffQuestion := models.Question{
		QuestionUuid: diffQuestionUuid,
		Title: title,
		Description: description,
	}
	err = Questions.UpsertQuestion(diffQuestion)
	if err != nil {
		t.Error("Unable to insert a different question")
	}

	diffInsertedQuestion, exists := Questions.SelectQuestion(diffQuestionUuid)
	if !exists {
		t.Error("Unable to insert a different question")
	}
	if diffInsertedQuestion != diffQuestion {
		t.Error("Unable to insert a different question correctly")
	}

	bothQuestions := Questions.SelectQuestions()
	if len(bothQuestions) != 2 {
		t.Error("Unable to keep two questions")
	}
}

func TestQuestionDaoImpl_DeleteQuestion(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	questions := Questions.SelectQuestions()
	if len(questions) != 0 {
		t.Error("Unable to start with empty database")
	}

	questionUuid := uuid.New()
	title := "Title"
	description := "Description"
	question := models.Question{
		QuestionUuid: questionUuid,
		Title: title,
		Description: description,
	}

	err := Questions.UpsertQuestion(question)
	if err != nil {
		t.Error("Unable to insert question")
	}

	questions = Questions.SelectQuestions()
	if len(questions) != 1 {
		t.Error("Unable to insert correctly")
	}

	err = Questions.DeleteQuestion(question.QuestionUuid)
	if err != nil {
		t.Error("Unable to delete question")
	}

	questions = Questions.SelectQuestions()
	if len(questions) != 0 {
		t.Error("Unable to delete question correctly")
	}
}