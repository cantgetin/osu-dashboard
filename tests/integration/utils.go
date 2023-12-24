package integration

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

func ClearTables(ctx context.Context, db *gorm.DB) error {
	name, ok := ctx.Value("environment").(string)
	if !ok {
		return fmt.Errorf("integration test variable environment not found in context")
	}

	if name != "integration-test" {
		return fmt.Errorf("not an integration test environment")
	}

	if err := db.Exec(`
	TRUNCATE TABLE users;
	TRUNCATE TABLE beatmaps;
	TRUNCATE TABLE mapsets;
`).Error; err != nil {
		return err
	}

	return nil
}
