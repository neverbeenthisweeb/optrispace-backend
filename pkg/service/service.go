package service

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"optrispace.com/work/pkg/model"
	"optrispace.com/work/pkg/service/pgsvc"
)

type (
	// GenericCRUD represents the standart methods for CRUD operations
	GenericCRUD[E any] interface {
		// Add saves the entity into storage
		Add(ctx context.Context, e *E) (*E, error)
		// Get reads specified entity from storage by specified id
		// It can return model.ErrNotFound
		Get(ctx context.Context, id string) (*E, error)

		// List reads all items from storage
		List(ctx context.Context) ([]*E, error)
	}

	// Security creates user representation from echo.Context, if there is such data
	Security interface {
		// FromEchoContext acquires user from echo and persistent storage
		// It will construct *model.UserContext in the context too
		FromEchoContext(c echo.Context) (*model.UserContext, error)

		// FromEchoContextByBasicAuth extracts user creds from basic auth header for specified realm
		FromEchoContextByBasicAuth(c echo.Context, realm string) (*model.UserContext, error)

		// FromLoginPassword creates UserContext from login and password in default realm
		FromLoginPassword(ctx context.Context, login, password string) (*model.UserContext, error)
	}

	// Job handles job offers
	Job interface {
		GenericCRUD[model.Job]

		// Patch partially updates existing Job object
		Patch(ctx context.Context, id, actorID string, patch map[string]any) (*model.Job, error)
	}

	// Person is a person who pay or earn
	Person interface {
		GenericCRUD[model.Person]

		// Update password
		UpdatePassword(ctx context.Context, subjectID, oldPassword, newPassword string) error

		// Patch partially updates existing Person object
		Patch(ctx context.Context, id, actorID string, patch map[string]any) error

		// SetResources fully replace resources for the person
		SetResources(ctx context.Context, id, actorID string, resources []byte) error
	}

	// Application is application for a job offer
	Application interface {
		GenericCRUD[model.Application]
		// ListBy returns list of entities by specified filters
		// If jobID != "", method returns list of jobs
		// if actorID != "", method returns list of applications for job author or applications
		ListBy(ctx context.Context, jobID, actorID string) ([]*model.Application, error)

		// ListByApplicant returns list of applications by specified applicant
		ListByApplicant(ctx context.Context, applicantID string) ([]*model.Application, error)
	}

	// Contract is an agreement between a Customer and a Performer (Contractor)
	Contract interface {
		// Add saves the entity into storage
		Add(ctx context.Context, c *model.Contract) (*model.Contract, error)

		// GetByIDForPerson reads specified entity from storage by specified id and related for person
		// It can return model.ErrNotFound
		GetByIDForPerson(ctx context.Context, id, personID string) (*model.Contract, error)

		// ListByPersonID returns list of entities by specific Person
		ListByPersonID(ctx context.Context, personID string) ([]*model.Contract, error)

		// Accept makes contract accepted if any
		Accept(ctx context.Context, id, actorID, performerAddress string) (*model.Contract, error)

		// Deploy makes contract accepted if any
		Deploy(ctx context.Context, id, actorID, contractAddress string) (*model.Contract, error)

		// Send makes contract sent if any
		Send(ctx context.Context, id, actorID string) (*model.Contract, error)

		// Approve makes contract approved if any
		Approve(ctx context.Context, id, actorID string) (*model.Contract, error)

		// Complete makes contract completed if any
		Complete(ctx context.Context, id, actorID string) (*model.Contract, error)
	}

	// Notification service manipulates with notifications
	Notification interface {
		// Push sends a data message to the configured channels (chats in Telegram messenger)
		Push(ctx context.Context, data string) error
	}

	// Stats service for statistic information
	Stats interface {
		Stats(ctx context.Context) (*model.Stats, error)
	}
)

// NewSecurity creates job service
func NewSecurity(db *sql.DB) Security {
	return pgsvc.NewSecurity(db)
}

// NewJob creates job service
func NewJob(db *sql.DB) Job {
	return pgsvc.NewJob(db)
}

// NewPerson creates person service
func NewPerson(db *sql.DB) Person {
	return pgsvc.NewPerson(db)
}

// NewApplication creates application service
func NewApplication(db *sql.DB) Application {
	return pgsvc.NewApplication(db)
}

// NewContract creates contract service
func NewContract(db *sql.DB) Contract {
	return pgsvc.NewContract(db)
}

// NewNotification creates notification service
func NewNotification(tgToken string, chatIDs ...int64) Notification {
	return pgsvc.NewNotification(tgToken, chatIDs...)
}

// NewStats create stats service
func NewStats(db *sql.DB) Stats {
	return pgsvc.NewStats(db)
}
