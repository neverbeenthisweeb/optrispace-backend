// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: messages.sql

package pgdao

import (
	"context"
	"time"
)

const messageAdd = `-- name: MessageAdd :one
insert into messages (
    id, chat_id, created_by, text
) values (
    $1, $2, $3, $4
) returning id, chat_id, created_at, created_by, text
`

type MessageAddParams struct {
	ID        string
	ChatID    string
	CreatedBy string
	Text      string
}

func (q *Queries) MessageAdd(ctx context.Context, arg MessageAddParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, messageAdd,
		arg.ID,
		arg.ChatID,
		arg.CreatedBy,
		arg.Text,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.Text,
	)
	return i, err
}

const messagesListByChat = `-- name: MessagesListByChat :many
select 
     m.id, m.chat_id, m.created_at, m.created_by, m.text
    ,p.display_name
from messages m 
    join persons p on m.created_by = p.id
where m.chat_id = $1::varchar
order by m.created_at asc
`

type MessagesListByChatRow struct {
	ID          string
	ChatID      string
	CreatedAt   time.Time
	CreatedBy   string
	Text        string
	DisplayName string
}

func (q *Queries) MessagesListByChat(ctx context.Context, chatID string) ([]MessagesListByChatRow, error) {
	rows, err := q.db.QueryContext(ctx, messagesListByChat, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MessagesListByChatRow
	for rows.Next() {
		var i MessagesListByChatRow
		if err := rows.Scan(
			&i.ID,
			&i.ChatID,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.Text,
			&i.DisplayName,
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

const messagesPurge = `-- name: MessagesPurge :exec
DELETE FROM messages
`

// Handle with care!
func (q *Queries) MessagesPurge(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, messagesPurge)
	return err
}