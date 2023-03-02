package security

import (
	"context"
	"errors"
	"github.com/tsingsun/woocoo/pkg/security"
	"strconv"
)

var (
	ErrInvalidUserID = errors.New("invalid user")
)

func UserIDFromContext(ctx context.Context) int {
	gp := security.GenericIdentityFromContext(ctx)
	id, err := strconv.Atoi(gp.Name())
	if err != nil || id == 0 {
		panic(ErrInvalidUserID)
	}
	return id
}
