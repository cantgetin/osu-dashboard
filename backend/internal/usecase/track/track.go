package track

import (
	"context"
	"database/sql"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/usecase/command"
	"playcount-monitor-backend/internal/usecase/mappers"
	"strconv"
)

func (uc *UseCase) Track(
	ctx context.Context,
) error {
	// get all following IDs from db and get updated data from api, update data in db
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

		for _, following := range follows {
			// get data from api
			user, err := uc.osuApiService.GetUser(ctx, strconv.Itoa(following.ID))
			if err != nil {
				return err
			}
			userMapsets, err := uc.osuApiService.GetUserMapsets(ctx, strconv.Itoa(following.ID))
			if err != nil {
				return err
			}

			// create/update data in db
			txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
				userExists, err := uc.user.Exists(ctx, tx, user.ID)
				if err != nil {
					return err
				}

				if userExists {
					// update user

				} else {
					// create usercard
					cmd := &command.CreateUserCardCommand{
						User:    mapOsuAPiUserToCreateUserCommand(user),
						Mapsets: mapOsuApiMapsetsToCreateMapsetCommands(userMapsets),
					}

					// create user
					userModel, err := mappers.MapCreateUserCardCommandToUserModel(cmd)
					if err != nil {
						return err
					}

					err = uc.user.Create(ctx, tx, userModel)
					if err != nil {
						return err
					}

					// create user mapsets
					for _, ms := range cmd.Mapsets {
						// create mapset
						var mapset *model.Mapset
						mapset, err = mappers.MapCreateMapsetCommandToMapsetModel(ms)
						if err != nil {
							return err
						}

						err = uc.mapset.Create(ctx, tx, mapset)
						if err != nil {
							return err
						}

						// create mapset beatmaps
						for _, bm := range ms.Beatmaps {
							var beatmap *model.Beatmap
							beatmap, err = mappers.MapCreateBeatmapCommandToBeatmapModel(bm)
							if err != nil {
								return err
							}

							err = uc.beatmap.Create(ctx, tx, beatmap)
						}
					}
				}

				return nil
			})
			if txErr != nil {
				return txErr
			}

		}

	}
}
