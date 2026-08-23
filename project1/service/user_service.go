package service

import (
    "fmt"
    "go-user-app/model"
    "go-user-app/repository"
)

type UserService struct {
    Repo repository.UserRepository
}

func (s UserService) FetchUser(id int) (model.User, error) {
    user, err := s.Repo.GetUser(id)
    if err != nil {
        return model.User{}, err
    }

    fmt.Println("Fetched user:", user.Name)
    return user, nil
}