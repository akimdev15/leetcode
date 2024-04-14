// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: questions.sql

package database

import (
	"context"
	"time"
)

const createQuestion = `-- name: CreateQuestion :one
INSERT INTO questions (id, name, url, solved, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, url, solved, updated_at
`

type CreateQuestionParams struct {
	ID        string
	Name      string
	Url       string
	Solved    string
	UpdatedAt time.Time
}

func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestion,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.Solved,
		arg.UpdatedAt,
	)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.Solved,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllQuestions = `-- name: GetAllQuestions :many
SELECT id, name, url, solved, updated_at FROM questions
`

func (q *Queries) GetAllQuestions(ctx context.Context) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, getAllQuestions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.Solved,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRandomQuestionURL = `-- name: GetRandomQuestionURL :one
SELECT url from questions ORDER BY RANDOM() LIMIT 1
`

func (q *Queries) GetRandomQuestionURL(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getRandomQuestionURL)
	var url string
	err := row.Scan(&url)
	return url, err
}
