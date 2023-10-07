package persistence

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"log"
	"strconv"
	"strings"
)

const (
	questionsTable = "questions"
	createQuestionsTable       = `
CREATE TABLE IF NOT EXISTS questions(
    question_uuid VARCHAR(36) NOT NULL PRIMARY KEY, 
    title VARCHAR(255) NOT NULL, 
    description VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL);`
	selectQuestionsTable       = `SELECT question_uuid, title, description, created_at from questions`
	upsertQuestionsTable = `
INSERT INTO questions(question_uuid, title, description, created_at) 
VALUES('{question_uuid}', '{title}', '{description}', '{created_at}') 
ON CONFLICT(question_uuid)
DO
UPDATE SET title='{title}', description='{description}', created_at='{created_at}';`
	deleteQuestionsTable = `DELETE FROM questions`
)

var Questions QuestionDao = QuestionDaoImpl{
	db: DB,
}

type QuestionDaoImpl struct {
	db *sql.DB
}

type QuestionDao interface {
	SelectQuestion(questionUuid uuid.UUID) (models.Question, bool)
	SelectQuestions() []models.Question
	UpsertQuestion(question models.Question) error
	DeleteQuestion(questionUuid uuid.UUID) error
}

func (q QuestionDaoImpl) SelectQuestion(questionUuid uuid.UUID) (models.Question, bool) {
	selectQuery := fmt.Sprintf("%s WHERE question_uuid='%s';", selectQuestionsTable, questionUuid)
	rows, err := q.db.Query(selectQuery)
	if err != nil {
		log.Println("Unable to select row", err)
		return models.Question{}, false
	}

	for rows.Next() {
		question := models.Question{}
		err := rows.Scan(&question.QuestionUuid, &question.Title, &question.Description, &question.CreatedAt)
		if err != nil {
			log.Println("Unable to scan row", err)
			return models.Question{}, false
		}
		return question, true
	}
	return models.Question{}, false
}

func (q QuestionDaoImpl) SelectQuestions() []models.Question {
	rows, err := q.db.Query(selectQuestionsTable)
	if err != nil {
		log.Println("Unable to select row", err)
		return []models.Question{}
	}

	var questions []models.Question
	for rows.Next() {
		question := models.Question{}
		err := rows.Scan(
			&question.QuestionUuid,
			&question.Title,
			&question.Description,
			&question.CreatedAt,
			)
		if err != nil {
			log.Println("Unable to scan row", err)
		}
		questions = append(questions, question)
	}
	return questions
}

func (q QuestionDaoImpl) UpsertQuestion(question models.Question) error {
	upsertQuery := upsertQuestionsTable
	upsertQuery = strings.ReplaceAll(upsertQuery, "{question_uuid}", question.QuestionUuid.String())
	upsertQuery = strings.ReplaceAll(upsertQuery, "{title}", question.Title)
	upsertQuery = strings.ReplaceAll(upsertQuery, "{description}", question.Description)
	upsertQuery = strings.ReplaceAll(upsertQuery, "{created_at}", strconv.FormatInt(question.CreatedAt, 10))
	_, err := q.db.Exec(upsertQuery)
	if err != nil {
		log.Println("Unable to upsert row", err)
		return err
	}
	return nil
}

func (q QuestionDaoImpl) DeleteQuestion(questionUuid uuid.UUID) error {
	_, err := q.db.Exec(fmt.Sprintf("%s WHERE question_uuid='%s';", deleteQuestionsTable, questionUuid))
	if err != nil {
		log.Println("Unable to delete row", err)
		return err
	}
	return nil
}
