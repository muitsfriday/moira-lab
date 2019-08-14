package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// User struct
type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Alias    string `json:"alias" form:"alias"`
}

// CreateUserBody body of create user request
type CreateUserBody struct {
	User User `json:"user" form:"user"`
}

// CreateUserResponse response for create user request
type CreateUserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func handleCreateUser(c echo.Context) error {
	b := new(CreateUserBody)
	if err := c.Bind(b); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, b.User)
}

func main() {
	e := echo.New()

	e.POST("/", handleCreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
