// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: persons.sql

package pgdao

import (
	"context"
	"database/sql"
	"encoding/json"
)

const personAdd = `-- name: PersonAdd :one
insert into persons (
    id, realm, login, password_hash, display_name, email, access_token, ethereum_address
) values (
    $1, $2, $3, $4, $5, $6, $7, $8
)
returning id, realm, login, password_hash, display_name, created_at, email, ethereum_address, resources, access_token, is_admin
`

type PersonAddParams struct {
	ID              string
	Realm           string
	Login           string
	PasswordHash    string
	DisplayName     string
	Email           string
	AccessToken     sql.NullString
	EthereumAddress string
}

func (q *Queries) PersonAdd(ctx context.Context, arg PersonAddParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, personAdd,
		arg.ID,
		arg.Realm,
		arg.Login,
		arg.PasswordHash,
		arg.DisplayName,
		arg.Email,
		arg.AccessToken,
		arg.EthereumAddress,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Realm,
		&i.Login,
		&i.PasswordHash,
		&i.DisplayName,
		&i.CreatedAt,
		&i.Email,
		&i.EthereumAddress,
		&i.Resources,
		&i.AccessToken,
		&i.IsAdmin,
	)
	return i, err
}

const personGet = `-- name: PersonGet :one
select id, realm, login, password_hash, display_name, created_at, email, ethereum_address, resources, access_token, is_admin from persons
	where id = $1::varchar
`

func (q *Queries) PersonGet(ctx context.Context, id string) (Person, error) {
	row := q.db.QueryRowContext(ctx, personGet, id)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Realm,
		&i.Login,
		&i.PasswordHash,
		&i.DisplayName,
		&i.CreatedAt,
		&i.Email,
		&i.EthereumAddress,
		&i.Resources,
		&i.AccessToken,
		&i.IsAdmin,
	)
	return i, err
}

const personGetByAccessToken = `-- name: PersonGetByAccessToken :one
select id, realm, login, password_hash, display_name, created_at, email, ethereum_address, resources, access_token, is_admin from persons
	where access_token = $1::varchar
`

func (q *Queries) PersonGetByAccessToken(ctx context.Context, accessToken string) (Person, error) {
	row := q.db.QueryRowContext(ctx, personGetByAccessToken, accessToken)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Realm,
		&i.Login,
		&i.PasswordHash,
		&i.DisplayName,
		&i.CreatedAt,
		&i.Email,
		&i.EthereumAddress,
		&i.Resources,
		&i.AccessToken,
		&i.IsAdmin,
	)
	return i, err
}

const personGetByLogin = `-- name: PersonGetByLogin :one
select id, realm, login, password_hash, display_name, created_at, email, ethereum_address, resources, access_token, is_admin from persons p
	where p.login = $1::varchar and p.realm = $2::varchar
`

type PersonGetByLoginParams struct {
	Login string
	Realm string
}

func (q *Queries) PersonGetByLogin(ctx context.Context, arg PersonGetByLoginParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, personGetByLogin, arg.Login, arg.Realm)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Realm,
		&i.Login,
		&i.PasswordHash,
		&i.DisplayName,
		&i.CreatedAt,
		&i.Email,
		&i.EthereumAddress,
		&i.Resources,
		&i.AccessToken,
		&i.IsAdmin,
	)
	return i, err
}

const personPatch = `-- name: PersonPatch :one
update persons
set
    ethereum_address = case when $1::boolean
        then $2::varchar else ethereum_address end

    , display_name = case when $3::boolean
        then $4::varchar else display_name end

where
    id = $5::varchar
returning id, realm, login, password_hash, display_name, created_at, email, ethereum_address, resources, access_token, is_admin
`

type PersonPatchParams struct {
	EthereumAddressChange bool
	EthereumAddress       string
	DisplayNameChange     bool
	DisplayName           string
	ID                    string
}

func (q *Queries) PersonPatch(ctx context.Context, arg PersonPatchParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, personPatch,
		arg.EthereumAddressChange,
		arg.EthereumAddress,
		arg.DisplayNameChange,
		arg.DisplayName,
		arg.ID,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Realm,
		&i.Login,
		&i.PasswordHash,
		&i.DisplayName,
		&i.CreatedAt,
		&i.Email,
		&i.EthereumAddress,
		&i.Resources,
		&i.AccessToken,
		&i.IsAdmin,
	)
	return i, err
}

const personSetAccessToken = `-- name: PersonSetAccessToken :exec
update persons
set
    access_token = $1::varchar
where
    id = $2::varchar
returning id, realm, login, password_hash, display_name, created_at, email, ethereum_address, resources, access_token, is_admin
`

type PersonSetAccessTokenParams struct {
	AccessToken string
	ID          string
}

// Sets the person's access token
func (q *Queries) PersonSetAccessToken(ctx context.Context, arg PersonSetAccessTokenParams) error {
	_, err := q.db.ExecContext(ctx, personSetAccessToken, arg.AccessToken, arg.ID)
	return err
}

const personSetEthereumAddress = `-- name: PersonSetEthereumAddress :exec
update persons
set
    ethereum_address = $1
where
    id = $2::varchar
`

type PersonSetEthereumAddressParams struct {
	EthereumAddress string
	ID              string
}

func (q *Queries) PersonSetEthereumAddress(ctx context.Context, arg PersonSetEthereumAddressParams) error {
	_, err := q.db.ExecContext(ctx, personSetEthereumAddress, arg.EthereumAddress, arg.ID)
	return err
}

const personSetIsAdmin = `-- name: PersonSetIsAdmin :exec
update persons
set
    is_admin = $1::boolean
where
    id = $2::varchar
`

type PersonSetIsAdminParams struct {
	IsAdmin bool
	ID      string
}

func (q *Queries) PersonSetIsAdmin(ctx context.Context, arg PersonSetIsAdminParams) error {
	_, err := q.db.ExecContext(ctx, personSetIsAdmin, arg.IsAdmin, arg.ID)
	return err
}

const personSetPassword = `-- name: PersonSetPassword :exec
update persons
set
    password_hash = $1::varchar
where
    id = $2::varchar
`

type PersonSetPasswordParams struct {
	NewPasswordHash string
	ID              string
}

func (q *Queries) PersonSetPassword(ctx context.Context, arg PersonSetPasswordParams) error {
	_, err := q.db.ExecContext(ctx, personSetPassword, arg.NewPasswordHash, arg.ID)
	return err
}

const personSetResources = `-- name: PersonSetResources :exec
update persons
set
    resources = $1::json
where
    id = $2::varchar
`

type PersonSetResourcesParams struct {
	Resources json.RawMessage
	ID        string
}

func (q *Queries) PersonSetResources(ctx context.Context, arg PersonSetResourcesParams) error {
	_, err := q.db.ExecContext(ctx, personSetResources, arg.Resources, arg.ID)
	return err
}

const personsList = `-- name: PersonsList :many
select id, realm, login, password_hash, display_name, created_at, email, ethereum_address, resources, access_token, is_admin from persons
`

func (q *Queries) PersonsList(ctx context.Context) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, personsList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.Realm,
			&i.Login,
			&i.PasswordHash,
			&i.DisplayName,
			&i.CreatedAt,
			&i.Email,
			&i.EthereumAddress,
			&i.Resources,
			&i.AccessToken,
			&i.IsAdmin,
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

const personsPurge = `-- name: PersonsPurge :exec
DELETE FROM persons
`

// Handle with care!
func (q *Queries) PersonsPurge(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, personsPurge)
	return err
}
