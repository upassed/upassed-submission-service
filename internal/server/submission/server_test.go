package submission_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-submission-service/internal/config"
	"github.com/upassed/upassed-submission-service/internal/handling"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-submission-service/internal/server"
	"github.com/upassed/upassed-submission-service/internal/util"
	"github.com/upassed/upassed-submission-service/internal/util/mocks"
	"github.com/upassed/upassed-submission-service/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	submissionClient client.SubmissionClient
	submissionSvc    *mocks.SubmissionService
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

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("cfg load error: ", err)
	}

	logger := logging.New(cfg.Env)
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	authClient := mocks.NewAuthClientMW(ctrl)
	authClient.EXPECT().AuthenticationUnaryServerInterceptor().Return(emptyAuthMiddleware())

	submissionSvc = mocks.NewSubmissionService(ctrl)
	submissionServer := server.New(server.AppServerCreateParams{
		Config:            cfg,
		Log:               logger,
		SubmissionService: submissionSvc,
		AuthClient:        authClient,
	})

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	cc, err := grpc.NewClient(fmt.Sprintf(":%s", cfg.GrpcServer.Port), opts...)
	if err != nil {
		log.Fatal("error creating client connection", err)
	}

	submissionClient = client.NewSubmissionClient(cc)
	go func() {
		if err := submissionServer.Run(); err != nil {
			os.Exit(1)
		}
	}()

	exitCode := m.Run()
	submissionServer.GracefulStop()
	os.Exit(exitCode)
}

func TestFindStudentFormSubmissions_InvalidRequest(t *testing.T) {
	request := &client.FindStudentFormSubmissionsRequest{
		FormId:          "invalid_uuid",
		StudentUsername: gofakeit.Username(),
	}

	_, err := submissionClient.FindStudentFormSubmissions(context.Background(), request)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, codes.InvalidArgument, convertedError.Code())
}

func TestFindStudentFormSubmissions_ServiceError(t *testing.T) {
	request := &client.FindStudentFormSubmissionsRequest{
		FormId:          uuid.NewString(),
		StudentUsername: gofakeit.Username(),
	}

	expectedServiceError := handling.New("some service error", codes.NotFound)
	submissionSvc.EXPECT().
		FindStudentFormSubmissions(gomock.Any(), gomock.Any()).
		Return(nil, expectedServiceError)

	_, err := submissionClient.FindStudentFormSubmissions(context.Background(), request)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedServiceError.Code(), convertedError.Code())
	assert.Equal(t, expectedServiceError.Error(), convertedError.Message())
}

func TestFindStudentFormSubmissions_HappyPath(t *testing.T) {
	request := &client.FindStudentFormSubmissionsRequest{
		FormId:          uuid.NewString(),
		StudentUsername: gofakeit.Username(),
	}

	foundFormSubmissions := util.RandomBusinessFormSubmissions()
	submissionSvc.EXPECT().
		FindStudentFormSubmissions(gomock.Any(), gomock.Any()).
		Return(foundFormSubmissions, nil)

	response, err := submissionClient.FindStudentFormSubmissions(context.Background(), request)
	require.NoError(t, err)

	assert.Equal(t, foundFormSubmissions.FormID.String(), response.GetFormId())
}

func emptyAuthMiddleware() func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctxWithUsername := context.WithValue(ctx, auth.UsernameKey, "someUsername")

		return handler(ctxWithUsername, req)
	}
}
