// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: visits.sql

package repository

import (
	"context"
	"time"
)

const countAllVisits = `-- name: CountAllVisits :one
SELECT COUNT(*) from visit
`

func (q *Queries) CountAllVisits(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countAllVisits)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countVisitors = `-- name: CountVisitors :one
SELECT COUNT(distinct(ip)) FROM visit WHERE visited_at > $1
`

func (q *Queries) CountVisitors(ctx context.Context, visitedAt time.Time) (int64, error) {
	row := q.db.QueryRow(ctx, countVisitors, visitedAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const insertVisit = `-- name: InsertVisit :exec
INSERT INTO visit (ip, visited_at) VALUES (($1::varchar)::inet, now())
`

func (q *Queries) InsertVisit(ctx context.Context, ip string) error {
	_, err := q.db.Exec(ctx, insertVisit, ip)
	return err
}
