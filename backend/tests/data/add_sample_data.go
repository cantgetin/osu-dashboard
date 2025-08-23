package main

import (
	"context"
	"github.com/caarlos0/env"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/model"
	"time"
)

type ContextKey string

const EnvKey ContextKey = "environment"

func main() {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to parse cfg, %v", err)
	}

	ctx := context.WithValue(context.Background(), EnvKey, "integration-test")

	if err := addSampleData(ctx, cfg); err != nil {
		log.Fatalf("failed to start app, %v", err)
	}
}

func addSampleData(ctx context.Context, cfg *config.Config) error {
	gdb, err := initDB(cfg)
	if err != nil {
		return err
	}

	user := &model.User{
		ID:        7192129,
		Username:  "Gasha",
		AvatarURL: "https://a.ppy.sh/7192129?1602378137.jpeg",
		UserStats: repository.JSON(`
{"2023-12-24T12:00:00Z":{"play_count":11000,"favourite_count":2, "map_count":3}, 
"2023-12-25T12:00:00Z":{"play_count":11090,"favourite_count":3, "map_count":4},
"2023-12-26T12:00:00Z":{"play_count":11200,"favourite_count":4, "map_count":5},
"2023-12-27T12:00:00Z":{"play_count":11634,"favourite_count":10, "map_count":6}}`),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	mapset := &model.Mapset{
		ID:     2015413,
		Artist: "Rizza",
		Title:  "bizzare",
		Covers: repository.JSON(`
{"cover": "https://assets.ppy.sh/beatmaps/2015413/covers/cover.jpg?1690122670",
"cover@2x": "https://assets.ppy.sh/beatmaps/2015413/covers/cover@2x.jpg?1690122670",
"card": "https://assets.ppy.sh/beatmaps/2015413/covers/card.jpg?1690122670",
"card@2x": "https://assets.ppy.sh/beatmaps/2015413/covers/card@2x.jpg?1690122670",
"list": "https://assets.ppy.sh/beatmaps/2015413/covers/list.jpg?1690122670",
"list@2x": "https://assets.ppy.sh/beatmaps/2015413/covers/list@2x.jpg?1690122670",
"slimcover": "https://assets.ppy.sh/beatmaps/2015413/covers/slimcover.jpg?1690122670",
"slimcover@2x": "https://assets.ppy.sh/beatmaps/2015413/covers/slimcover@2x.jpg?1690122670"}`),
		Status:      "graveyard",
		LastUpdated: time.Now().UTC(),
		UserID:      7192129,
		Creator:     "Gasha",
		PreviewURL:  "//b.ppy.sh/preview/2015413.mp3",
		Tags:        "rap trap hyperpop synthwave chill girl rizza sqwore",
		BPM:         150,
		MapsetStats: repository.JSON(`
{"2023-12-24T12:00:00Z":{"play_count":654,"favourite_count":2},
"2023-12-25T12:00:00Z":{"play_count":800,"favourite_count":3},
"2023-12-26T12:00:00Z":{"play_count":2000,"favourite_count":4},
"2023-12-27T12:00:00Z":{"play_count":2300,"favourite_count":15}}`),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	beatmaps := []model.Beatmap{
		{
			ID:               4195095,
			MapsetID:         2015413,
			DifficultyRating: 5.63,
			Version:          "diff1",
			Accuracy:         8.4,
			AR:               9.3,
			BPM:              150,
			CS:               3.9,
			Status:           "graveyard",
			URL:              "https://osu.ppy.sh/beatmaps/4195095",
			TotalLength:      114,
			UserID:           7192129,
			LastUpdated:      time.Now().UTC(),
			BeatmapStats: repository.JSON(`
{"2023-12-24T12:00:00Z":{"play_count":10,"pass_count":5},
"2023-12-25T12:00:00Z":{"play_count":20,"pass_count":10},
"2023-12-26T12:00:00Z":{"play_count":100,"pass_count":15},
"2023-12-27T12:00:00Z":{"play_count":340,"pass_count":20}}`),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		{
			ID:               4195096,
			MapsetID:         2015413,
			DifficultyRating: 5.67,
			Version:          "diff2",
			Accuracy:         8.6,
			AR:               9.2,
			BPM:              150,
			CS:               4,
			Status:           "graveyard",
			URL:              "https://osu.ppy.sh/beatmaps/4195096",
			TotalLength:      115,
			UserID:           7192129,
			LastUpdated:      time.Now().UTC(),
			BeatmapStats: repository.JSON(`
{"2023-12-24T12:00:00Z":{"play_count":10,"pass_count":5},
"2023-12-25T12:00:00Z":{"play_count":20,"pass_count":10},
"2023-12-26T12:00:00Z":{"play_count":90,"pass_count":15},
"2023-12-27T12:00:00Z":{"play_count":260,"pass_count":20}}`),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}

	err = gdb.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}

	err = gdb.Create(&mapset).Error
	if err != nil {
		return err
	}

	for _, beatmap := range beatmaps {
		err = gdb.Create(&beatmap).Error
		if err != nil {
			return err
		}
	}

	db, err := gdb.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Error(err)
	}

	log.Infof("sample data added")
	return nil
}

func initDB(cfg *config.Config) (*gorm.DB, error) {
	cfg.PgDSN = "postgresql://localhost:5432/db?user=db&password=db&sslmode=disable"
	var gdb *gorm.DB

	var err error
	gdb, err = bootstrap.InitDB(cfg)
	if err != nil {
		return nil, err
	}
	if _, err = gdb.DB(); err != nil {
		return nil, err
	}

	db, err := gdb.DB()
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join(dir, "./", "migrations"),
	}

	migrationsCount, err := countMigrations("./migrations")
	if err != nil {
		log.Fatalf("Error counting migrations: %v", err)
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Infof("executed %d migrations", n)
		log.Fatalf("Could not run the 'UP' migrations: %v", err)
	} else {
		if n < migrationsCount {
			log.Fatalf("should be at least %v migrations", migrationsCount)
		}
		log.Infof("executed %d migrations", n)
	}

	return gdb, nil
}

func countMigrations(folderPath string) (int, error) {
	files, err := filepath.Glob(filepath.Join(folderPath, "*"))
	if err != nil {
		return 0, err
	}
	return len(files), nil
}
