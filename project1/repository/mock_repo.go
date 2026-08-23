package repository

import "go-user-app/model"

type MockRepo struct{}

func (m MockRepo) GetUser(id int) (model.User, error) {
    return model.User{
        ID: id,
        Name: "Mock User",
        Email: "mock@test.com",
    }, nil
}