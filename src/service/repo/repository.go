package repo

import (
	"database/sql"
	"fmt"
)

type Repository[T any] struct {
	DB *sql.DB // hoặc GORM, hoặc gì đó
}

func (r *Repository[T]) Create(item T) error {
	fmt.Println("Base Create")
	return nil
}

func (r *Repository[T]) GetByID(id string) (T, error) {
	fmt.Println("Base GetByID")
	var result T
	return result, nil
}

func (r *Repository[T]) Update(item T) error {
	fmt.Println("Base Update")
	return nil
}

func (r *Repository[T]) Delete(id int) error {
	fmt.Println("Base Delete")
	return nil
}
