package repository

import "go-user-app/model"

type PostgresRepo struct{}

func (p PostgresRepo) GetUser(id int) (model.User, error) {
    return model.User{
        ID: id,
        Name: "Postgres User",
        Email: "pg@test.com",
    }, nil
}