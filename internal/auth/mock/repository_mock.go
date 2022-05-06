package mock

import (
	"context"

	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/models"
)

type MockRepository struct {
}

func NewMockRepository() *MockRepository {
	mock := &MockRepository{}
	return mock
}

func (m *MockRepository) Register(ctx context.Context, user *models.User) error {
	return nil
}
