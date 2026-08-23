package handler

import (
    "fmt"
    "go-user-app/service"
)

type UserHandler struct {
    Service service.UserService
}

func (h UserHandler) GetUser(id int) {
    user, err := h.Service.FetchUser(id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("User Response:", user)
}