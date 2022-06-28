package localwt

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"

	auth2 "github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/venus-auth/auth"
	"github.com/filecoin-project/venus-auth/core"
	"github.com/gbrlsnchs/jwt/v3"
)

type LocalJwt struct {
	alg jwt.Algorithm
}

func NewLocalJwt() (*LocalJwt, error) {
	sk, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 32))
	if err != nil {
		return nil, err
	}
	return &LocalJwt{alg: jwt.NewHS256(sk)}, nil
}

func (localJwt *LocalJwt) Token() ([]byte, error) {
	payload := auth.JWTPayload{
		Name:  "chain-co-admin",
		Perm:  "admin",
		Extra: "",
	}

	return jwt.Sign(&payload, localJwt.alg)
}

func (localJwt *LocalJwt) Verify(ctx context.Context, token string) ([]auth2.Permission, error) {
	var payload auth.JWTPayload
	if _, err := jwt.Verify([]byte(token), localJwt.alg, &payload); err != nil {
		return nil, fmt.Errorf("JWT Verification failed: %w", err)
	}

	jwtPerms := core.AdaptOldStrategy(payload.Perm)
	perms := make([]auth2.Permission, len(jwtPerms))
	copy(perms, jwtPerms)
	return perms, nil
}
