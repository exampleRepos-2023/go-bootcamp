package persistence

import (
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"testing"
)

func TestQuestionAnswersDaoImpl_SelectQuestionAnswersByQuestionUuid(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	questionUuid := uuid.New()
	answerUuid := uuid.New()
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: questionUuid,
		AnswerUuid: answerUuid,
	}
	questionAnswersS := QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(questionAnswersS) != 0 {
		t.Error("Unable to start with empty database")
	}

	err := QuestionAnswersS.InsertQuestionAnswers(questionAnswers)
	if err != nil {
		t.Error("Unable to upsert question")
	}

	questionAnswersS = QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(questionAnswersS) != 1 {
		t.Error("Unable to insert one questionAnswers")
	}
	if questionAnswersS[0] != questionAnswers {
		t.Error("Unable to insert question correctly")
	}
}

func TestQuestionAnswersDaoImpl_InsertQuestionAnswers(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	questionUuid := uuid.New()
	answerUuid := uuid.New()
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: questionUuid,
		AnswerUuid: answerUuid,
	}

	questionAnswersS := QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(questionAnswersS) != 0 {
		t.Error("Unable to start with empty database")
	}

	err := QuestionAnswersS.InsertQuestionAnswers(questionAnswers)
	if err != nil {
		t.Error("Unable to insert questionAnswers")
	}

	insertedQuestionAnswersS := QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(insertedQuestionAnswersS) != 1 {
		t.Error("Unable to insert questionAnswers")
	}
	if insertedQuestionAnswersS[0] != questionAnswers {
		t.Error("Unable to insert questionAnswers correctly")
	}

	diffAnswerUuid := uuid.New()
	sameQuestionQuestionAnswers := models.QuestionAnswers{
		QuestionUuid: questionUuid,
		AnswerUuid: diffAnswerUuid,
	}
	err = QuestionAnswersS.InsertQuestionAnswers(sameQuestionQuestionAnswers)
	if err != nil {
		t.Error("Unable to insert a different answer questionAnswers")
	}

	sameQuestionQuestionAnswersS := QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionUuid)
	if len(sameQuestionQuestionAnswersS) != 2 {
		t.Error("Unable to insert a different questionAnswers")
	}
	if sameQuestionQuestionAnswersS[0] != sameQuestionQuestionAnswers &&
		sameQuestionQuestionAnswersS[1] != sameQuestionQuestionAnswers {
		t.Error("Unable to insert a different questionAnswers correctly")
	}
}

func TestQuestionAnswersDaoImpl_DeleteQuestionAnswers(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	questionUuid := uuid.New()
	answerUuid := uuid.New()
	questionAnswers := models.QuestionAnswers{
		QuestionUuid: questionUuid,
		AnswerUuid: answerUuid,
	}
	diffAnswerUuid1 := uuid.New()
	sameQuestionQuestionAnswers1 := models.QuestionAnswers{
		QuestionUuid: questionUuid,
		AnswerUuid: diffAnswerUuid1,
	}
	diffAnswerUuid2 := uuid.New()
	sameQuestionQuestionAnswers2 := models.QuestionAnswers{
		QuestionUuid: questionUuid,
		AnswerUuid: diffAnswerUuid2,
	}

	questionAnswersS := QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(questionAnswersS) != 0 {
		t.Error("Unable to start with empty database")
	}

	err := QuestionAnswersS.InsertQuestionAnswers(questionAnswers)
	if err != nil {
		t.Error("Unable to insert questionAnswers")
	}
	err = QuestionAnswersS.InsertQuestionAnswers(sameQuestionQuestionAnswers1)
	if err != nil {
		t.Error("Unable to insert same question questionAnswers 1")
	}
	err = QuestionAnswersS.InsertQuestionAnswers(sameQuestionQuestionAnswers2)
	if err != nil {
		t.Error("Unable to insert same question questionAnswers 2")
	}

	insertedQuestionAnswersS := QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(insertedQuestionAnswersS) != 3 {
		t.Error("Unable to insert questionAnswers")
	}

	err = QuestionAnswersS.DeleteQuestionAnswersByAnswerUuid(questionAnswers.AnswerUuid)
	if err != nil {
		t.Error("Unable to delete questionAnswers by answer")
	}
	insertedQuestionAnswersS = QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(insertedQuestionAnswersS) != 2 {
		t.Error("Unable to delete questionAnswers by answer correctly")
	}

	err = QuestionAnswersS.DeleteQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if err != nil {
		t.Error("Unable to delete questionAnswers by question")
	}
	insertedQuestionAnswersS = QuestionAnswersS.SelectQuestionAnswersByQuestionUuid(questionAnswers.QuestionUuid)
	if len(insertedQuestionAnswersS) != 0 {
		t.Error("Unable to delete questionAnswers by question correctly")
	}
}