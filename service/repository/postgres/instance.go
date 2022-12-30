package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type handle struct {
	Conn *sql.DB
}

// ConnectPostgreSQL
// @docs https://pkg.go.dev/github.com/lib/pq
// dsn e.g.
// postgres://username:password@localhost:5432/mydb?sslmode=verify-full
// postgres://username:password@localhost:5432/mydb?sslmode=disable
func ConnectPostgreSQL(dsn string) (*handle, error) {

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		conn.Close()
		return nil, err
	}
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(20)

	return &handle{
		Conn: conn,
	}, nil
}
