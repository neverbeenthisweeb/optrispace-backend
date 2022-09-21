// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: contracts.sql

package pgdao

import (
	"context"
	"database/sql"
	"time"
)

const contractAdd = `-- name: ContractAdd :one
insert into contracts (
    id, customer_id, performer_id, application_id, title, description, price, duration, created_by, customer_address, performer_address, status, contract_address
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
returning id, customer_id, performer_id, application_id, title, description, price, duration, status, created_by, created_at, updated_at, customer_address, performer_address, contract_address
`

type ContractAddParams struct {
	ID               string
	CustomerID       string
	PerformerID      string
	ApplicationID    string
	Title            string
	Description      string
	Price            string
	Duration         sql.NullInt32
	CreatedBy        string
	CustomerAddress  string
	PerformerAddress string
	Status           string
	ContractAddress  string
}

func (q *Queries) ContractAdd(ctx context.Context, arg ContractAddParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, contractAdd,
		arg.ID,
		arg.CustomerID,
		arg.PerformerID,
		arg.ApplicationID,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.Duration,
		arg.CreatedBy,
		arg.CustomerAddress,
		arg.PerformerAddress,
		arg.Status,
		arg.ContractAddress,
	)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.PerformerID,
		&i.ApplicationID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.Duration,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CustomerAddress,
		&i.PerformerAddress,
		&i.ContractAddress,
	)
	return i, err
}

const contractGet = `-- name: ContractGet :one
select c.id, c.customer_id, c.performer_id, c.application_id, c.title, c.description, c.price, c.duration, c.status, c.created_by, c.created_at, c.updated_at, c.customer_address, c.performer_address, c.contract_address from contracts c
join applications a on a.id = c.application_id and a.applicant_id = c.performer_id
join jobs j on j.id = a.job_id
join persons customer on customer.id = c.customer_id
join persons performer on performer.id = c.performer_id
where c.id = $1::varchar
`

// mostly in testing purposes
func (q *Queries) ContractGet(ctx context.Context, id string) (Contract, error) {
	row := q.db.QueryRowContext(ctx, contractGet, id)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.PerformerID,
		&i.ApplicationID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.Duration,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CustomerAddress,
		&i.PerformerAddress,
		&i.ContractAddress,
	)
	return i, err
}

const contractGetByIDAndPersonID = `-- name: ContractGetByIDAndPersonID :one
select c.id, c.customer_id, c.performer_id, c.application_id, c.title, c.description, c.price, c.duration, c.status, c.created_by, c.created_at, c.updated_at, c.customer_address, c.performer_address, c.contract_address from contracts c
join applications a on a.id = c.application_id and a.applicant_id = c.performer_id
join jobs j on j.id = a.job_id
join persons customer on customer.id = c.customer_id
join persons performer on performer.id = c.performer_id
where c.id = $1::varchar and (c.customer_id = $2::varchar or c.performer_id = $2::varchar)
`

type ContractGetByIDAndPersonIDParams struct {
	ID       string
	PersonID string
}

func (q *Queries) ContractGetByIDAndPersonID(ctx context.Context, arg ContractGetByIDAndPersonIDParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, contractGetByIDAndPersonID, arg.ID, arg.PersonID)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.PerformerID,
		&i.ApplicationID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.Duration,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CustomerAddress,
		&i.PerformerAddress,
		&i.ContractAddress,
	)
	return i, err
}

const contractPatch = `-- name: ContractPatch :one
update contracts
set
    status = case when $1::boolean
        then $2::varchar else status end,

    performer_address = case when $3::boolean
        then $4::varchar else performer_address end,

    contract_address = case when $5::boolean
        then $6::varchar else contract_address end
where
    id = $7::varchar
returning id, customer_id, performer_id, application_id, title, description, price, duration, status, created_by, created_at, updated_at, customer_address, performer_address, contract_address
`

type ContractPatchParams struct {
	StatusChange           bool
	Status                 string
	PerformerAddressChange bool
	PerformerAddress       string
	ContractAddressChange  bool
	ContractAddress        string
	ID                     string
}

func (q *Queries) ContractPatch(ctx context.Context, arg ContractPatchParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, contractPatch,
		arg.StatusChange,
		arg.Status,
		arg.PerformerAddressChange,
		arg.PerformerAddress,
		arg.ContractAddressChange,
		arg.ContractAddress,
		arg.ID,
	)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.PerformerID,
		&i.ApplicationID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.Duration,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CustomerAddress,
		&i.PerformerAddress,
		&i.ContractAddress,
	)
	return i, err
}

const contractsGetByPerson = `-- name: ContractsGetByPerson :many
select
    c.id, c.customer_id, c.performer_id, c.application_id, c.title, c.description, c.price, c.duration, c.status, c.created_by, c.created_at, c.updated_at, c.customer_address, c.performer_address, c.contract_address
    , pc.display_name as customer_name
    , pp.display_name as performer_name
from contracts c
left join persons pc on c.customer_id = pc.id
left join persons pp on c.performer_id = pp.id
where c.customer_id = $1::varchar or c.performer_id = $1::varchar
order by c.created_at desc
`

type ContractsGetByPersonRow struct {
	ID               string
	CustomerID       string
	PerformerID      string
	ApplicationID    string
	Title            string
	Description      string
	Price            string
	Duration         sql.NullInt32
	Status           string
	CreatedBy        string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CustomerAddress  string
	PerformerAddress string
	ContractAddress  string
	CustomerName     sql.NullString
	PerformerName    sql.NullString
}

func (q *Queries) ContractsGetByPerson(ctx context.Context, personID string) ([]ContractsGetByPersonRow, error) {
	rows, err := q.db.QueryContext(ctx, contractsGetByPerson, personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ContractsGetByPersonRow
	for rows.Next() {
		var i ContractsGetByPersonRow
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.PerformerID,
			&i.ApplicationID,
			&i.Title,
			&i.Description,
			&i.Price,
			&i.Duration,
			&i.Status,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CustomerAddress,
			&i.PerformerAddress,
			&i.ContractAddress,
			&i.CustomerName,
			&i.PerformerName,
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

const contractsPurge = `-- name: ContractsPurge :exec
DELETE FROM contracts
`

// Handle with care!
func (q *Queries) ContractsPurge(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, contractsPurge)
	return err
}
