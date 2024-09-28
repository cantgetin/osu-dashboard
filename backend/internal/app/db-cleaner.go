package app

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/app/dbcleaner"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/cleanrepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/usecase/cleaner"
	"time"
)

func RunDBCleaner(
	ctx context.Context,
	cfg *config.Config,
	lg *log.Logger,
) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	err = bootstrap.ApplyMigrations(db)
	if err != nil {
		return err
	}

	const waitForConnection = 5 * time.Second
	txm := bootstrap.ConnectTxManager("db-cleaner", waitForConnection, db, lg)

	// init repos
	userRepo := userrepository.New(cfg, lg)
	mapsetRepo := mapsetrepository.New(cfg, lg)
	beatmapRepo := beatmaprepository.New(cfg, lg)
	cleanerRepo := cleanrepository.New(cfg, lg)

	// init usecase
	cleanerUc := cleaner.New(cfg, lg, txm, userRepo, mapsetRepo, beatmapRepo, cleanerRepo)

	c := dbcleaner.New(cfg, lg, cleanerUc)

	c.Start(ctx)

	gracefulShutDown(ctx, cancel)

	return nil
}
