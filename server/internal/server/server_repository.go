package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"server/util"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type respository struct {
	db DBTX
}

func NewRepository(db DBTX) *respository {
	return &respository{db: db}
}

func (r *respository) CreateServer(ctx context.Context, server *CreateServerReq, c *gin.Context) (*CreateServerRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 插入 server
	query := "INSERT INTO server (server_name) VALUES (?)"
	row, err := tx.ExecContext(ctx, query, server.ServerName)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("LastInsertId. ", err)
		return nil, err
	}

	// We parse the JWT token from the cookie to get logging user's ID
	jwtClaims, err := util.ParseJWT(c)
	if err != nil {
		tx.Rollback()
		fmt.Println("ParseJWT error: ", err)
		return nil, err
	}
	
	user_id:= jwtClaims["id"].(string)

	// 插入 joins
	query = "INSERT INTO joins (user_id, server_id) VALUES (?, ?)"
	_, err = tx.ExecContext(ctx, query, user_id, lastInsertID)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql err. ", err)
		return nil, err
	}

	// 插入 owns
	query = "INSERT INTO owns (user_id, server_id) VALUES (?, ?)"
	_, err = tx.ExecContext(ctx, query, user_id, lastInsertID)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql err. ", err)
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("tx.Commit. ", err)
		return nil, err
	}

	res := &CreateServerRes{
		ServerId:   int(lastInsertID),
		ServerName: server.ServerName,
	}

	return res, nil
}

// func (r *respository) GetServerByEmail(ctx context.Context, email string) (*User, error) {
// 	s := GetServerRes
// 	query := "SELECT user_id, email, user_name, password FROM user WHERE email = ?"
// 	if err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password); err != nil {
// 		return &User{}, err
// 	}

// 	return &u, nil
// }