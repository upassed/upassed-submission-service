package submission_test

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-submission-service/internal/config"
	"github.com/upassed/upassed-submission-service/internal/handling"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
	"github.com/upassed/upassed-submission-service/internal/service/submission"
	"github.com/upassed/upassed-submission-service/internal/util"
	"github.com/upassed/upassed-submission-service/internal/util/mocks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	cfg        *config.Config
	repository *mocks.SubmissionRepository
	service    submission.Service
)

func TestMain(m *testing.M) {
	currentDir, _ := os.Getwd()
	projectRoot, err := util.GetProjectRoot(currentDir)
	if err != nil {
		log.Fatal("error to get project root folder: ", err)
	}

	if err := os.Setenv(config.EnvConfigPath, filepath.Join(projectRoot, "config", "test.yml")); err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Load()
	if err != nil {
		log.Fatal("unable to parse config: ", err)
	}

	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	repository = mocks.NewSubmissionRepository(ctrl)
	service = submission.New(cfg, logging.New(config.EnvTesting), repository)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreate_ErrorCheckingExistingSubmissions(t *testing.T) {
	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	submissionToCreate := util.RandomBusinessSubmission()
	expectedRepositoryError := errors.New("some repo error")

	repository.EXPECT().
		Exists(gomock.Any(), gomock.Any()).
		Return(false, expectedRepositoryError)

	_, err := service.Create(ctx, submissionToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepositoryError.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestCreate_ErrorDeletingExistingSubmissions(t *testing.T) {
	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	submissionToCreate := util.RandomBusinessSubmission()
	repository.EXPECT().
		Exists(gomock.Any(), gomock.Any()).
		Return(true, nil)

	expectedRepositoryError := errors.New("some repo error")
	repository.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(expectedRepositoryError)

	_, err := service.Create(ctx, submissionToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepositoryError.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestCreate_ErrorSavingSubmissions(t *testing.T) {
	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	submissionToCreate := util.RandomBusinessSubmission()
	repository.EXPECT().
		Exists(gomock.Any(), gomock.Any()).
		Return(false, nil)

	expectedRepositoryError := errors.New("some repo error")
	repository.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(expectedRepositoryError)

	_, err := service.Create(ctx, submissionToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepositoryError.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestCreate_ErrorDeadlineExceeded(t *testing.T) {
	oldTimeout := cfg.Timeouts.EndpointExecutionTimeoutMS
	cfg.Timeouts.EndpointExecutionTimeoutMS = "0"

	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	submissionToCreate := util.RandomBusinessSubmission()
	repository.EXPECT().
		Exists(gomock.Any(), gomock.Any()).
		Return(false, nil)

	repository.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil)

	_, err := service.Create(ctx, submissionToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, submission.ErrSubmissionCreateDeadlineExceeded.Error(), convertedError.Message())
	assert.Equal(t, codes.DeadlineExceeded, convertedError.Code())

	cfg.Timeouts.EndpointExecutionTimeoutMS = oldTimeout
}

func TestCreate_HappyPath_SubmissionsExisted(t *testing.T) {
	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	submissionToCreate := util.RandomBusinessSubmission()
	repository.EXPECT().
		Exists(gomock.Any(), gomock.Any()).
		Return(true, nil)

	repository.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	repository.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil)

	response, err := service.Create(ctx, submissionToCreate)
	require.NoError(t, err)

	assert.Equal(t, len(submissionToCreate.AnswerIDs), len(response.CreatedSubmissionIDs))
}

func TestCreate_HappyPath_NoSubmissionsExisted(t *testing.T) {
	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	submissionToCreate := util.RandomBusinessSubmission()
	repository.EXPECT().
		Exists(gomock.Any(), gomock.Any()).
		Return(false, nil)

	repository.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil)

	response, err := service.Create(ctx, submissionToCreate)
	require.NoError(t, err)

	assert.Equal(t, len(submissionToCreate.AnswerIDs), len(response.CreatedSubmissionIDs))
}

func TestFindStudentFormSubmissions_RepositoryError(t *testing.T) {
	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	expectedRepositoryError := handling.New("repo error", codes.NotFound)
	repository.EXPECT().
		FindStudentFormSubmissions(gomock.Any(), gomock.Any()).
		Return(nil, expectedRepositoryError)

	formID := uuid.New()
	_, err := service.FindStudentFormSubmissions(ctx, &business.StudentFormSubmissionSearchParams{
		StudentUsername: studentUsername,
		FormID:          formID,
	})

	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepositoryError.Error(), convertedError.Message())
	assert.Equal(t, expectedRepositoryError.Code(), convertedError.Code())
}

func TestFindStudentFormSubmissions_ErrorDeadlineExceeded(t *testing.T) {
	oldTimeout := cfg.Timeouts.EndpointExecutionTimeoutMS
	cfg.Timeouts.EndpointExecutionTimeoutMS = "0"

	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	foundSubmissions := util.RandomDomainSubmissions()
	repository.EXPECT().
		FindStudentFormSubmissions(gomock.Any(), gomock.Any()).
		Return(foundSubmissions, nil)

	formID := uuid.New()
	_, err := service.FindStudentFormSubmissions(ctx, &business.StudentFormSubmissionSearchParams{
		StudentUsername: studentUsername,
		FormID:          formID,
	})

	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, submission.ErrSubmissionSearchingDeadlineExceeded.Error(), convertedError.Message())
	assert.Equal(t, codes.DeadlineExceeded, convertedError.Code())

	cfg.Timeouts.EndpointExecutionTimeoutMS = oldTimeout
}

func TestFindStudentFormSubmissions_HappyPath(t *testing.T) {
	studentUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, studentUsername)

	foundSubmissions := util.RandomDomainSubmissions()
	repository.EXPECT().
		FindStudentFormSubmissions(gomock.Any(), gomock.Any()).
		Return(foundSubmissions, nil)

	formID := uuid.New()
	response, err := service.FindStudentFormSubmissions(ctx, &business.StudentFormSubmissionSearchParams{
		StudentUsername: studentUsername,
		FormID:          formID,
	})

	require.NoError(t, err)
	assert.NotNil(t, response.StudentUsername)
	assert.NotNil(t, response.FormID)
	assert.NotEmpty(t, response.QuestionSubmissions)
}
