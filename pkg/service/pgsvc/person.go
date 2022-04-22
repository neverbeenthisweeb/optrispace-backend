package pgsvc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jaswdr/faker"
	"optrispace.com/work/pkg/db/pgdao"
	"optrispace.com/work/pkg/model"
)

type (
	// PersonSvc is a person service
	PersonSvc struct {
		db *sql.DB
	}
)

// NewPerson creates service
func NewPerson(db *sql.DB) *PersonSvc {
	return &PersonSvc{db: db}
}

// Add implements service.Person
func (s *PersonSvc) Add(ctx context.Context, person *model.Person) (*model.Person, error) {
	var result *model.Person
	return result, doWithQueries(ctx, s.db, defaultRwTxOpts, func(queries *pgdao.Queries) error {
		input := pgdao.PersonAddParams{
			ID:           pgdao.NewID(),
			Realm:        person.Realm,
			Login:        person.Login,
			PasswordHash: CreateHashFromPassword(person.Password),
			DisplayName:  person.DisplayName,
			Email:        person.Email,
		}

		if input.Realm == "" {
			input.Realm = model.InhouseRealm
		}

		f := faker.New()

		if input.Login == "" {
			input.Login = input.ID
		}

		if input.DisplayName == "" {
			input.DisplayName = f.Person().Name()
		}

		p, err := queries.PersonAdd(ctx, input)
		if err != nil {
			return fmt.Errorf("unable to PersonAdd: %w", err)
		}

		result = &model.Person{
			ID:          p.ID,
			Realm:       p.Realm,
			Login:       p.Login,
			DisplayName: p.DisplayName,
			CreatedAt:   p.CreatedAt,
			Email:       p.Email,
		}

		return nil
	})
}

// Get implements service.Person
func (s *PersonSvc) Get(ctx context.Context, id string) (*model.Person, error) {
	var result *model.Person
	return result, doWithQueries(ctx, s.db, defaultRoTxOpts, func(queries *pgdao.Queries) error {
		p, err := queries.PersonGet(ctx, id)

		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrEntityNotFound
		}

		if err != nil {
			return fmt.Errorf("unable to PersonGet: %w", err)
		}

		result = &model.Person{
			ID:          p.ID,
			Realm:       p.Realm,
			Login:       p.Login,
			DisplayName: p.DisplayName,
			CreatedAt:   p.CreatedAt,
			Email:       p.Email,
		}

		return nil
	})
}

// List implements service.Person
func (s *PersonSvc) List(ctx context.Context) ([]*model.Person, error) {
	result := make([]*model.Person, 0)
	return result, doWithQueries(ctx, s.db, defaultRoTxOpts, func(queries *pgdao.Queries) error {
		pp, err := queries.PersonsList(ctx)
		if err != nil {
			return err
		}

		for _, p := range pp {
			result = append(result, &model.Person{
				ID:          p.ID,
				Realm:       p.Realm,
				Login:       p.Login,
				DisplayName: p.DisplayName,
				CreatedAt:   p.CreatedAt,
				Email:       p.Email,
			})
		}

		return nil
	})
}