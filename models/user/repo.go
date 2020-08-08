package user

import (
	"context"
)

type Collection interface {
	Insert(ctx context.Context, user *User) error
	Get(ctx context.Context, username, password, guid, refreshToken, accessToken string) (*User, error)
}