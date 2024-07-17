package server

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
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type respository struct {
	db DBTX
}

func NewRepository(db DBTX) *respository {
	return &respository{db: db}
}

func (r *respository) CreateServer(ctx context.Context, req *CreateServerReq, creator int) (*ServerMetadata, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 插入 server
	query := "INSERT INTO server (server_name) VALUES (?)"
	row, err := tx.ExecContext(ctx, query, req.ServerName)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("LastInsertId. ", err)
		return nil, err
	}

	// fmt.Printf("lastInsertID: %d, uid: %d\n", lastInsertID, creator)

	// 插入 joins
	query = "INSERT INTO joins (user_id, server_id) VALUES (?, ?)"
	_, err = tx.ExecContext(ctx, query, creator, lastInsertID)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql err. ", err)
		return nil, err
	}

	// 插入 owns
	query = "INSERT INTO owns (user_id, server_id) VALUES (?, ?)"
	_, err = tx.ExecContext(ctx, query, creator, lastInsertID)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql err. ", err)
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("tx.Commit. ", err)
		return nil, err
	}

	res := &ServerMetadata{
		ServerId:   int(lastInsertID),
		ServerName: req.ServerName,
	}

	return res, nil
}

func (r *respository) GetServerByEmail(ctx context.Context, email string) (*GetServerRes, error) {
	query := "SELECT server_id, server_name FROM user NATURAL JOIN server NATURAL JOIN joins WHERE email = ?"
	rows, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	servers := make(map[string]string)

	for rows.Next() {
		var serverID string
		var serverName string
		if err := rows.Scan(&serverID, &serverName); err != nil {
			return nil, err
		}

		servers[serverID] = serverName
	}

	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	err = rows.Close()
	if err != nil {
		fmt.Println("rows.Close. ", err)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err = rows.Err(); err != nil {
		fmt.Println("rows.Err. ", err)
	}

	return &GetServerRes{Servers: servers}, nil
}
