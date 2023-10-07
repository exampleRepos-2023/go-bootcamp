package persistence

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"log"
	"strings"
)

const (
	questionAnswersTable = "question_answers"
	createQuestionAnswersTable = `
CREATE TABLE IF NOT EXISTS question_answers(
    question_uuid VARCHAR(36) NOT NULL, 
    answer_uuid VARCHAR(36) NOT NULL UNIQUE);`
	selectQuestionAnswersTable = `SELECT question_uuid, answer_uuid from question_answers`
	insertQuestionAnswersTable = `
INSERT INTO question_answers(question_uuid, answer_uuid) 
VALUES('{question_uuid}', '{answer_uuid}')`
	deleteQuestionAnswersTable = `
DELETE FROM question_answers`
)

var QuestionAnswersS QuestionAnswersDao = QuestionAnswersDaoImpl{
	db: DB,
}

type QuestionAnswersDaoImpl struct {
	db *sql.DB
}

type QuestionAnswersDao interface {
	SelectQuestionAnswersByQuestionUuid(questionUuid uuid.UUID) []models.QuestionAnswers
	InsertQuestionAnswers(questionAnswers models.QuestionAnswers) error
	DeleteQuestionAnswersByQuestionUuid(questionUuid uuid.UUID) error
	DeleteQuestionAnswersByAnswerUuid(answerUuid uuid.UUID) error
}

func (qa QuestionAnswersDaoImpl) SelectQuestionAnswersByQuestionUuid(questionUuid uuid.UUID) []models.QuestionAnswers {
	selectQuery := fmt.Sprintf("%s WHERE question_uuid='%s';", selectQuestionAnswersTable, questionUuid)
	rows, err := DB.Query(selectQuery)
	if err != nil {
		log.Println("Unable to select row", err)
		return []models.QuestionAnswers{}
	}

	var questionAnswersS []models.QuestionAnswers
	for rows.Next() {
		questionAnswers := models.QuestionAnswers{}
		err := rows.Scan(&questionAnswers.QuestionUuid, &questionAnswers.AnswerUuid)
		if err != nil {
			log.Println("Unable to scan row", err)
			return []models.QuestionAnswers{}
		}
		questionAnswersS = append(questionAnswersS, questionAnswers)
	}
	return questionAnswersS
}

func (qa QuestionAnswersDaoImpl) InsertQuestionAnswers(questionAnswers models.QuestionAnswers) error {
	upsertQuery := insertQuestionAnswersTable
	upsertQuery = strings.ReplaceAll(upsertQuery, "{question_uuid}", questionAnswers.QuestionUuid.String())
	upsertQuery = strings.ReplaceAll(upsertQuery, "{answer_uuid}", questionAnswers.AnswerUuid.String())
	_, err := DB.Exec(upsertQuery)
	if err != nil {
		log.Println("Unable to upsert row", err)
		return err
	}
	return nil
}

func (qa QuestionAnswersDaoImpl) DeleteQuestionAnswersByQuestionUuid(questionUuid uuid.UUID) error {
	_, err := DB.Exec(fmt.Sprintf("%s WHERE question_uuid='%s';", deleteQuestionAnswersTable, questionUuid))
	if err != nil {
		log.Println("Unable to delete row", err)
		return err
	}
	return nil
}

func (qa QuestionAnswersDaoImpl) DeleteQuestionAnswersByAnswerUuid(answerUuid uuid.UUID) error {
	_, err := DB.Exec(fmt.Sprintf("%s WHERE answer_uuid='%s';", deleteQuestionAnswersTable, answerUuid))
	if err != nil {
		log.Println("Unable to delete row", err)
		return err
	}
	return nil
}

func init() {
	QuestionAnswersS = QuestionAnswersDaoImpl{}
}
