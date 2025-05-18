package repo

import (
	"context"
	"fmt"
	"gin_example/src/model"
)

type UserRepo struct {
	*Repository[model.User] // kế thừa method CRUD
}

// Hàm mở rộng riêng
func (r *UserRepo) FindByName(ctx context.Context, name string) ([]model.User, error) {
	fmt.Println("FindByName:", name)
	return []model.User{}, nil
}

// Interface nếu cần đa hình
type IUser interface {
	IRepository[model.User]
	FindByName(ctx context.Context, name string) ([]model.User, error)
}
