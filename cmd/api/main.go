package main

import (
	"fmt"
	"log"
	"news/config"
	"news/internal/infra/repositories"
	"news/internal/usecase"
	"news/pkg/postgres"
	r "news/pkg/redis"
	"news/pkg/utils"
	"os"

	internalHttp "news/internal/infra/http"

	"github.com/go-redis/redis/v8"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	db, err := postgres.NewPgDB(
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.SSLMode,
		cfg.Postgres.Loc,
	)
	if err != nil {
		log.Fatalf("open db connection failed %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisClient.Host,
		Password: cfg.RedisClient.Password,
	})

	address := fmt.Sprintf(":%d", cfg.App.Port)

	redis := r.New(redisClient)
	newsMdl := repositories.NewNewsRepository(db)
	newsUc := usecase.NewNewsUsecase(newsMdl, redis)
	internalHttp.NewProcess(cfg, address, newsUc)

}
