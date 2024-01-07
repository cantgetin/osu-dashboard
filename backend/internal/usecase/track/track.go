package track

import (
	"context"
	"database/sql"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

func (uc *UseCase) Track(
	ctx context.Context,
) error {
	// get all following IDs and get data from api
	for {
		// get IDs
		var follows []*model.Following
		if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
			var err error
			if follows, err = uc.following.List(ctx, tx); err != nil {
				return err
			}
			return nil
		}, txmanager.Level(sql.LevelReadCommitted)); err != nil {
			return err
		}

		if len(follows) == 0 {
			return fmt.Errorf("no following users present in db")
		}

		for _, f := range follows {
			// get data from api
			_, err := uc.getCurrentUserCardFromAPI(ctx, f.ID)
			if err != nil {
				return err
			}

			// update userCardInfo
			// TODO
		}

	}
}

type WIPType struct {
	// TODO
}

func (uc *UseCase) getCurrentUserCardFromAPI(ctx context.Context, userID int) (*WIPType, error) {
	// TODO
	return nil, nil
}
