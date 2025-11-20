package bootstrap

import (
	"log"
	"os"

	"github.com/yasinahlattci/sre-case-study/app/internal/config"
	"github.com/yasinahlattci/sre-case-study/app/internal/db"
	"github.com/yasinahlattci/sre-case-study/app/internal/service"
)

type Dependencies struct {
	Logger  *log.Logger
	Service *service.DynamoDBService
	Config  *config.Config
}

func Bootstrap(env string) (*Dependencies, error) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(".conf", env)

	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
		return nil, err
	}

	client := db.GetClient(cfg.Database.Region)
	svc := service.NewDynamoDBService(client, cfg.Database.TableName)

	return &Dependencies{
		Logger:  logger,
		Service: svc,
		Config:  cfg,
	}, nil
}
