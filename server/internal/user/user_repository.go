package user

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type respository struct {
	db DBTX
}

func NewRepository(db DBTX) *respository {
	return &respository{db: db}
}

func (r *respository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO user (user_name, email, password) VALUES (?, ?, ?)"

	row, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		fmt.Println("db.QueryRowContext. ", err)
	}

	lastInsertID, err := row.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId. ", err)
	}

	user.ID = int(lastInsertID)
	return user, nil
}

func (r *respository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT user_id, email, user_name, password FROM user WHERE email = ?"
	if err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password); err != nil {
		fmt.Println("Error when get user by email. ", err)
		return &User{}, nil
	}

	return &u, nil
}
