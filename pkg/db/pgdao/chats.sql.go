// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: chats.sql

package pgdao

import (
	"context"
)

const chatAdd = `-- name: ChatAdd :one
insert into chats (
    id, topic
) values (
    $1, $2
) returning id, topic, created_at
`

type ChatAddParams struct {
	ID    string
	Topic string
}

func (q *Queries) ChatAdd(ctx context.Context, arg ChatAddParams) (Chat, error) {
	row := q.db.QueryRowContext(ctx, chatAdd, arg.ID, arg.Topic)
	var i Chat
	err := row.Scan(&i.ID, &i.Topic, &i.CreatedAt)
	return i, err
}

const chatGet = `-- name: ChatGet :one
select
    c.id, c.topic, c.created_at
from chats c
where c.id = $1::varchar
`

func (q *Queries) ChatGet(ctx context.Context, id string) (Chat, error) {
	row := q.db.QueryRowContext(ctx, chatGet, id)
	var i Chat
	err := row.Scan(&i.ID, &i.Topic, &i.CreatedAt)
	return i, err
}

const chatGetByTopic = `-- name: ChatGetByTopic :one
select
    c.id, c.topic, c.created_at
from chats c
where c.topic = $1::varchar
`

func (q *Queries) ChatGetByTopic(ctx context.Context, topic string) (Chat, error) {
	row := q.db.QueryRowContext(ctx, chatGetByTopic, topic)
	var i Chat
	err := row.Scan(&i.ID, &i.Topic, &i.CreatedAt)
	return i, err
}

const chatParticipantAdd = `-- name: ChatParticipantAdd :one
insert into chats_participants (
    chat_id, person_id
) values (
    $1, $2
) returning chat_id, person_id
`

type ChatParticipantAddParams struct {
	ChatID   string
	PersonID string
}

func (q *Queries) ChatParticipantAdd(ctx context.Context, arg ChatParticipantAddParams) (ChatsParticipant, error) {
	row := q.db.QueryRowContext(ctx, chatParticipantAdd, arg.ChatID, arg.PersonID)
	var i ChatsParticipant
	err := row.Scan(&i.ChatID, &i.PersonID)
	return i, err
}

const chatParticipantGet = `-- name: ChatParticipantGet :one
select
    cp.chat_id, cp.person_id
from chats_participants cp
where
    cp.chat_id = $1::varchar
    and cp.person_id = $2::varchar
`

type ChatParticipantGetParams struct {
	ChatID   string
	PersonID string
}

func (q *Queries) ChatParticipantGet(ctx context.Context, arg ChatParticipantGetParams) (ChatsParticipant, error) {
	row := q.db.QueryRowContext(ctx, chatParticipantGet, arg.ChatID, arg.PersonID)
	var i ChatsParticipant
	err := row.Scan(&i.ChatID, &i.PersonID)
	return i, err
}

const chatsPurge = `-- name: ChatsPurge :exec
DELETE FROM chats
`

// Handle with care!
func (q *Queries) ChatsPurge(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, chatsPurge)
	return err
}