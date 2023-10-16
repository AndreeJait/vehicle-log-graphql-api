package impl

import "context"

func (h handler) LogInUser(ctx context.Context, vehicleType, platNumber, townCode string) error {
	return h.ucLog.LogInUser(ctx, vehicleType, platNumber, townCode)
}
