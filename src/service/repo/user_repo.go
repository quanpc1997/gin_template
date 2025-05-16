package repo

import (
	"fmt"
	"gin_example/src/model"
)

type UserRepo struct {
	*Repository[model.User] // kế thừa method CRUD
}

// Hàm mở rộng riêng
func (r *UserRepo) FindByName(name string) ([]model.User, error) {
	fmt.Println("FindByName:", name)
	return []model.User{}, nil
}

// Interface nếu cần đa hình
type UserRepository interface {
	Create(model.User) error
	GetByID(string) (model.User, error)
	Update(model.User) error
	Delete(int) error
	FindByName(string) ([]model.User, error)
}
