package repository

import "go-user-app/model"

type UserRepository interface {
    GetUser(id int) (model.User, error)
}