package authenticator

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/models/user"
	"time"
)

type Authenticator struct {
	users      user.Users
	salt       string
	signingKey []byte
	exp        time.Duration
}

func (a *Authenticator) SignUp(ctx context.Context, user *user.User) error {
	user.Password = generateDBPwd(user, a.salt)
	return a.users.Insert(ctx, user)
}

func (a *Authenticator) SignIn(ctx context.Context, user *user.User) (string, error) {
	user.Password = generateDBPwd(user, a.salt)

	_, err := a.users.Get(ctx, user.Username, user.Password)
	if err != nil {
		return "", err
	}

	return generateToken(a.exp, user.Username, a.signingKey)
}

func generateDBPwd(user *user.User, salt string) string {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(salt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}

func generateToken(expireDuration time.Duration, username string, signingKey []byte) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(expireDuration),
			"IssuedAt":  time.Now(),
			"Username":  username,
		}).SignedString(signingKey)
}

func NewAuthenticator(coll user.Users, salt string, signingKey []byte, exp time.Duration) *Authenticator {
	return &Authenticator{
		users:      coll,
		salt:       salt,
		signingKey: signingKey,
		exp:        exp,
	}
}