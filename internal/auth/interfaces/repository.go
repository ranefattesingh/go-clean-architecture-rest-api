package interfaces

import (
	"context"

	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/models"
)

type RepositoryInterface interface {
	Register(context.Context, *models.User) error
}
