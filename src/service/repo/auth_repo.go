package repo

import (
	"context"
	"gin_example/src/model"
)

type AuthRepo struct {
	*Repository[model.User]
}

func (a *AuthRepo) Login(ctx context.Context) string {
	return "Ok"
}

type IAuthRepo interface {
}
