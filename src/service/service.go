package service

import (
	"fmt"
	"gin_example/src/service/repo"
)

type Service[T any] struct {
	Repo repo.Repository[T]
}

func (s *Service[T]) Create(item T) {
	err := s.Repo.Create(item)
	if err != nil {
		fmt.Println("Error creating user:", err)
	}
}

func (s *Service[T]) GetByID(id string) (T, error) {
	obj, err := s.Repo.GetByID(id)
	if err != nil {
		fmt.Println("Error geting by id service:", err)
	}
	return obj, err
}

func (s *Service[T]) Update(item T) error {
	err := s.Repo.Update(item)
	if err != nil {
		fmt.Println("Error updating user:", err)
	}
	return err
}

func (s *Service[T]) Delete(id int) error {
	err := s.Repo.Delete(id)
	if err != nil {
		fmt.Println("Error updating user:", err)
	}
	return err
}
