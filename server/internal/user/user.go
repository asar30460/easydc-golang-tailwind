package user

import (
	"context"
)

type User struct {
	// The suffixes after the types in the struct fields are called struct tags.
	// Struct tags are a form of metadata that can be attached to struct fields in Go.
	// They are used to provide additional information or instructions to libraries and tools that work with structs

	// These tags are used by libraries such as the encoding/json package for JSON operations

	ID       int    `json:"id" db:"id"`             // JSON key: "id", Database column: "id"
	Username string `json:"user_name" db:"user_name"` // JSON key: "username", Database column: "username"
	Email    string `json:"email" db:"email"`       // JSON key: "email", Database column: "email"
	Password string `json:"password" db:"password"` // JSON key: "password", Database column: "password"
}

type CreateUserReq struct {
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	accessToken string
	Email    string `json:"email"`
	Username    string `json:"user_name"`
}

type Respository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login (ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
}
