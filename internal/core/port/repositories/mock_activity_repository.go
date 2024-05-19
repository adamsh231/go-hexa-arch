package repositories

import (
	"go-hexa/internal/core/domain/models"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockActivityRepository struct {
	mock.Mock
}

func (m *MockActivityRepository) InsertActivity(activity models.ActivityModel) error {
	args := m.Called(activity)
	return args.Error(0)
}

func (m *MockActivityRepository) SearchActivities(service string, created time.Time, page, limit int64) (int64, []models.ActivityModel, error) {
	args := m.Called(service, created, page, limit)
	return int64(args.Int(0)), args.Get(1).([]models.ActivityModel), args.Error(2)
}

func (m *MockActivityRepository) FindActivity(id string) (models.ActivityModel, error) {
	args := m.Called(id)
	return args.Get(0).(models.ActivityModel), args.Error(1)
}
