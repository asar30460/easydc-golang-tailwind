package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"
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

func (r *respository) CreateServer(ctx context.Context, server_name string, creator int) (int, string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, "", err
	}

	// 插入 server
	query := "INSERT INTO server (server_name) VALUES (?)"
	row, err := tx.ExecContext(ctx, query, server_name)
	if err != nil {
		return -1, "", err
	}
	lastInsertID, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("LastInsertId. ", err)
		return -1, "", err
	}

	// fmt.Printf("lastInsertID: %d, uid: %d\n", lastInsertID, creator)

	// 插入 joins
	query = "INSERT INTO joins (user_id, server_id) VALUES (?, ?)"
	_, err = tx.ExecContext(ctx, query, creator, lastInsertID)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql err. ", err)
		return -1, "", err
	}

	// 插入 owns
	query = "INSERT INTO owns (user_id, server_id) VALUES (?, ?)"
	_, err = tx.ExecContext(ctx, query, creator, lastInsertID)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql err. ", err)
		return -1, "", err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("tx.Commit. ", err)
		return -1, "", err
	}

	return int(lastInsertID), server_name, nil
}

func (r *respository) GetServerByEmail(ctx context.Context, email string) (map[string]string, error) {
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

	return servers, nil
}

func (r *respository) CreateChannel(ctx context.Context, channel_name string, server_id int) (int, string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, "", err
	}

	// 插入 channel
	fmt.Println("channel_name: ", channel_name)
	query := "INSERT INTO channel (channel_name) VALUES (?)"
	row, err := tx.ExecContext(ctx, query, channel_name)
	if err != nil {
		return -1, "", err
	}
	lastInsertID, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("LastInsertId. ", err)
		return -1, "", err
	}

	// 插入 has
	query = "INSERT INTO has (server_id, channel_id) VALUES (?, ?)"
	_, err = tx.ExecContext(ctx, query, server_id, lastInsertID)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql err. ", err)
		return -1, "", err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("tx.Commit. ", err)
		return -1, "", err
	}

	return int(lastInsertID), channel_name, err
}

func (r *respository) GetChannel(ctx context.Context, server_id int) (map[string]string, error) {
	query := "SELECT channel_id, channel_name FROM channel NATURAL JOIN has WHERE server_id = ?"

	rows, err := r.db.QueryContext(ctx, query, server_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	channels := make(map[string]string)

	for rows.Next() {
		var channelID string
		var channelName string
		if err := rows.Scan(&channelID, &channelName); err != nil {
			return nil, err
		}

		channels[channelID] = channelName
	}

	err = rows.Close()
	if err != nil {
		fmt.Println("rows.Close. ", err)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("rows.Err. ", err)
	}

	return channels, nil
}

func (r *respository) GetMember(ctx context.Context, server_id int) (map[string]string, error) {
	query := "SELECT email, user_name FROM user NATURAL JOIN joins WHERE server_id = ?"

	rows, err := r.db.QueryContext(ctx, query, server_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make(map[string]string)

	for rows.Next() {
		var email string
		var userName string
		if err := rows.Scan(&email, &userName); err != nil {
			return nil, err
		}

		members[email] = userName
	}

	err = rows.Close()
	if err != nil {
		fmt.Println("rows.Close. ", err)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("rows.Err. ", err)
	}

	return members, nil
}

func (r *respository) CreateMsg(ctx context.Context, channel_id int, user_id int, time time.Time, message string) (Msg, error) {
	query := "INSERT INTO chat (channel_id, user_id, time, content) VALUES (?, ?, ?, ?)"
	row, err := r.db.QueryContext(ctx, query, channel_id, user_id, time, message)
	if err != nil {
		return Msg{
			UserID:   -1,
			UserName: "",
			Time:     time,
			Message:  "",
		}, err
	}
	defer row.Close()

	return Msg{UserID: user_id, UserName: "Not necessary here!", Time: time, Message: message}, err
}

func (r *respository) GetHistorysMsg(ctx context.Context, channel_id int) ([]Msg, error) {
	query := "SELECT user_id, user_name, time, content FROM channel NATURAL JOIN chat NATURAL JOIN user WHERE channel_id = ?"

	rows, err := r.db.QueryContext(ctx, query, channel_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []Msg
	for rows.Next() {
		var userID int
		var userName string
		var time time.Time
		var message string
		if err := rows.Scan(&userID, &userName, &time, &message); err != nil {
			return nil, err
		}
		msgs = append(msgs, Msg{UserID: userID, UserName: userName, Time: time, Message: message})
	}

	err = rows.Close()
	if err != nil {
		fmt.Println("rows.Close. ", err)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("rows.Err. ", err)
	}

	return msgs, nil
}
