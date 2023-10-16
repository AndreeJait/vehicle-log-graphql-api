package impl

import (
	"context"
	"vehicle-log-graphql-api/internal/common/publisher"
	"vehicle-log-graphql-api/internal/shared/ierr"
)

func (u useCase) LogInUser(ctx context.Context, vehicleType, platNumber, townCode string) error {
	_, available, err := u.redis.CheckAndGetTownAvailable(ctx, townCode, vehicleType)
	if err != nil {
		return err
	}
	_, quantity, err := u.redis.CheckAndGetTownQuantity(ctx, townCode, vehicleType)
	if err != nil {
		return err
	}

	if (available + 1) >= quantity {
		return ierr.ErrFullQuotaParkir
	}

	err = u.pubs.LogIn(ctx, publisher.MessageLogIn{
		TownCode:    townCode,
		PlatNumber:  platNumber,
		VehicleType: vehicleType,
	})

	return err
}
