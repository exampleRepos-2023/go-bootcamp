package persistence

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/3advanced/4finalproject/solution/models"
	"log"
	"strconv"
	"strings"
)

const (
	answersTable = "answers"
	createAnswersTable       = `
CREATE TABLE IF NOT EXISTS answers(
    answer_uuid VARCHAR(36) NOT NULL PRIMARY KEY, 
    content VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL);`
	selectAnswersTable       = `SELECT answer_uuid, content, created_at from answers`
	upsertAnswersTable = `
INSERT INTO answers(answer_uuid, content, created_at) 
VALUES('{answer_uuid}', '{content}', '{created_at}') 
ON CONFLICT(answer_uuid)
DO
UPDATE SET content='{content}', created_at='{created_at}';`
	deleteAnswersTable = `DELETE FROM answers`
)

var Answers AnswerDao = AnswerDaoImpl{}

type AnswerDaoImpl struct {

}

type AnswerDao interface {
	SelectAnswer(answerUuid uuid.UUID) (models.Answer, bool)
	SelectAnswers(answerUuids []uuid.UUID) []models.Answer
	UpsertAnswer(answer models.Answer) error
	DeleteAnswer(answerUuid uuid.UUID) error
	DeleteAnswers(answerUuids []uuid.UUID) error
}

func (a AnswerDaoImpl) SelectAnswer(answerUuid uuid.UUID) (models.Answer, bool) {
	selectQuery := fmt.Sprintf("%s WHERE answer_uuid='%s';", selectAnswersTable, answerUuid)
	rows, err := DB.Query(selectQuery)
	if err != nil {
		log.Println("Unable to select row", err)
		return models.Answer{}, false
	}

	for rows.Next() {
		answer := models.Answer{}
		err := rows.Scan(&answer.AnswerUuid, &answer.Content, &answer.CreatedAt)
		if err != nil {
			log.Println("Unable to scan row", err)
			return models.Answer{}, false
		}
		return answer, true
	}
	return models.Answer{}, false
}

func (a AnswerDaoImpl) SelectAnswers(answerUuids []uuid.UUID) []models.Answer {
	if len(answerUuids) == 0 {
		return []models.Answer{}
	}

	answerUuidsStringBuilder := strings.Builder{}
	for _, answerUuid := range answerUuids {
		answerUuidsStringBuilder.WriteString(fmt.Sprintf("'%s',",answerUuid.String()))
	}
	answerUuidsString := answerUuidsStringBuilder.String()
	answerUuidsString = answerUuidsString[:len(answerUuidsString)- 1]
	selectQuery := fmt.Sprintf("%s WHERE answer_uuid IN (%s);", selectAnswersTable, answerUuidsString)
	rows, err := DB.Query(selectQuery)
	if err != nil {
		log.Println("Unable to select row", err)
		return []models.Answer{}
	}

	var answers []models.Answer
	for rows.Next() {
		answer := models.Answer{}
		err := rows.Scan(
			&answer.AnswerUuid,
			&answer.Content,
			&answer.CreatedAt,
		)
		if err != nil {
			log.Println("Unable to scan row", err)
		}
		answers = append(answers, answer)
	}
	return answers
}

func (a AnswerDaoImpl) UpsertAnswer(answer models.Answer) error {
	upsertQuery := upsertAnswersTable
	upsertQuery = strings.ReplaceAll(upsertQuery, "{answer_uuid}", answer.AnswerUuid.String())
	upsertQuery = strings.ReplaceAll(upsertQuery, "{content}", answer.Content)
	upsertQuery = strings.ReplaceAll(upsertQuery, "{created_at}", strconv.FormatInt(answer.CreatedAt, 10))
	_, err := DB.Exec(upsertQuery)
	if err != nil {
		log.Println("Unable to upsert row", err)
		return err
	}
	return nil
}

func (a AnswerDaoImpl) DeleteAnswer(answerUuid uuid.UUID) error {
	return a.DeleteAnswers([]uuid.UUID{answerUuid})
}

func (a AnswerDaoImpl) DeleteAnswers(answerUuids []uuid.UUID) error {
	if len(answerUuids) == 0 {
		return nil
	}

	answerUuidsStringBuilder := strings.Builder{}
	for _, answerUuid := range answerUuids {
		answerUuidsStringBuilder.WriteString(fmt.Sprintf("'%s',",answerUuid.String()))
	}
	answerUuidsString := answerUuidsStringBuilder.String()
	answerUuidsString = answerUuidsString[:len(answerUuidsString) - 1]
	deleteQuery := fmt.Sprintf("%s WHERE answer_uuid IN (%s);", deleteAnswersTable, answerUuidsString)
	_, err := DB.Exec(deleteQuery)
	if err != nil {
		log.Println("Unable to delete row", err)
		return err
	}
	return nil
}
