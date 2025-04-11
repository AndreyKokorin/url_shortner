package service_test

import (
	"URL_shortner/internal/model"
	service "URL_shortner/internal/service/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (mc *mockRepository) NewUser(user *model.User) error {
	args := mc.Called(user)
	return args.Error(1)
}

func (mc *mockRepository) GetByEmail(email string) (*model.User, error) {
	args := mc.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestGetByEmail(t *testing.T) {
	mockRepo := new(mockRepository)

	fakePassword, _ := bcrypt.GenerateFromPassword([]byte("test_password"), bcrypt.DefaultCost)

	testUser := &model.User{Id: 0, Email: "test@test.com", Password: string(fakePassword)}
	mockRepo.On("GetByEmail", "test@test.com").Return(testUser, nil)

	userTestService := service.NewUserService(mockRepo)
	token, err := userTestService.LogIn("test@test.com", "test_password")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
