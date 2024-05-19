package services

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models"
	"go-hexa/internal/core/port/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertActivity(t *testing.T) {

	// init
	repo := new(repositories.MockActivityRepository)
	service := NewActivityService(repo)
	input := entities.InsertActivityInput{
		Service:   "test-service",
		Version:   "1.0",
		Message:   "test message",
		Activity:  "test activity",
		CreatedBy: "test-user",
		Data:      nil,
		Response:  nil,
	}

	// repo
	repo.On("InsertActivity", mock.Anything).Return(nil)
	err := service.InsertActivity(input)
	assert.NoError(t, err)

}

func TestSearchActivities(t *testing.T) {

	// init
	repo := new(repositories.MockActivityRepository)
	service := NewActivityService(repo)
	input := entities.SearchActivityInput{
		Service: "test-service",
		Created: time.Now(),
		Page:    1,
		Limit:   10,
	}

	// expected
	expectedActivities := []models.ActivityModel{
		{
			ID:        "1",
			Service:   "test-service",
			Version:   "1.0",
			Message:   "test message",
			Activity:  "test activity",
			CreatedBy: "test-user",
			Data:      nil,
			Response:  nil,
			Created:   time.Now(),
		},
	}

	// repo
	repo.On("SearchActivities", input.Service, input.Created, int64(input.Page), int64(input.Limit)).Return(len(expectedActivities), expectedActivities, nil)
	total, activities, err := service.SearchActivities(input)
	assert.NoError(t, err)

	// test
	assert.Equal(t, len(expectedActivities), total)
	assert.Equal(t, len(expectedActivities), len(activities))
}

func TestFindActivityByID(t *testing.T) {

	// init mock
	repo := new(repositories.MockActivityRepository)
	service := NewActivityService(repo)

	// expected
	expectedActivity := models.ActivityModel{
		ID:        "1",
		Service:   "test-service",
		Version:   "1.0",
		Message:   "test message",
		Activity:  "test activity",
		CreatedBy: "test-user",
		Data:      nil,
		Response:  nil,
		Created:   time.Now(),
	}

	// repo
	repo.On("FindActivity", "1").Return(expectedActivity, nil)
	output, err := service.FindActivityByID("1")
	assert.NoError(t, err)

	// test
	assert.Equal(t, expectedActivity.ID, output.ID)
	assert.Equal(t, expectedActivity.Service, output.Service)
	assert.Equal(t, expectedActivity.Version, output.Version)
	assert.Equal(t, expectedActivity.Message, output.Message)
	assert.Equal(t, expectedActivity.Activity, output.Activity)
	assert.Equal(t, expectedActivity.CreatedBy, output.CreatedBy)
	assert.Equal(t, expectedActivity.Data, output.Data)
	assert.Equal(t, expectedActivity.Response, output.Response)
	assert.Equal(t, expectedActivity.Created.Format(time.DateTime), output.Created)

}
