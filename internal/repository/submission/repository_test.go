package submission_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-submission-service/internal/config"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-submission-service/internal/repository"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
	"github.com/upassed/upassed-submission-service/internal/repository/submission"
	"github.com/upassed/upassed-submission-service/internal/testcontainer"
	"github.com/upassed/upassed-submission-service/internal/util"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

var (
	submissionRepository submission.Repository
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
		log.Fatal("unable to parse config: ", err)
	}

	ctx := context.Background()
	postgresTestcontainer, err := testcontainer.NewPostgresTestcontainer(ctx)
	if err != nil {
		log.Fatal("unable to create a testcontainer: ", err)
	}

	port, err := postgresTestcontainer.Start(ctx)
	if err != nil {
		log.Fatal("unable to get a postgres testcontainer real port: ", err)
	}

	cfg.Storage.Port = strconv.Itoa(port)
	logger := logging.New(cfg.Env)
	if err := postgresTestcontainer.Migrate(cfg, logger); err != nil {
		log.Fatal("unable to run migrations: ", err)
	}

	db, err := repository.OpenGormDbConnection(cfg, logger)
	if err != nil {
		log.Fatal("unable to open connection to postgres: ", err)
	}

	submissionRepository = submission.New(db, cfg, logger)
	exitCode := m.Run()
	if err := postgresTestcontainer.Stop(ctx); err != nil {
		log.Println("unable to stop postgres testcontainer: ", err)
	}

	os.Exit(exitCode)
}

func TestSaveAndExist_HappyPath(t *testing.T) {
	domainSubmissions := util.RandomDomainSubmissions()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, domainSubmissions[0].StudentUsername)

	exists, err := submissionRepository.Exists(ctx, &domain.SubmissionExistCheckParams{
		StudentUsername: domainSubmissions[0].StudentUsername,
		FormID:          domainSubmissions[0].FormID,
		QuestionID:      domainSubmissions[0].QuestionID,
	})

	require.NoError(t, err)
	assert.False(t, exists)

	err = submissionRepository.Save(ctx, domainSubmissions)
	require.NoError(t, err)

	exists, err = submissionRepository.Exists(ctx, &domain.SubmissionExistCheckParams{
		StudentUsername: domainSubmissions[0].StudentUsername,
		FormID:          domainSubmissions[0].FormID,
		QuestionID:      domainSubmissions[0].QuestionID,
	})

	require.NoError(t, err)
	assert.True(t, exists)
}
