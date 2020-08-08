package authenticator

import (
	"context"
	"server/models/user"
)

type Authenticable interface {
	SignUp(ctx context.Context, user *user.User) error
	SignIn(ctx context.Context, user *user.User) (string, error)
}
