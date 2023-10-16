package publisher

import "context"

type Publisher interface {
	LogIn(ctx context.Context, req MessageLogIn) error
	LogOut(ctx context.Context, req MessageLogOut) error
}
