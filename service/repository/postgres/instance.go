package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var _handle *handle

type handle struct {
	Conn *sql.DB
}

func getHandle() (*handle, error) {

	if _handle != nil {
		return _handle, nil
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	name := os.Getenv("POSTGRES_NAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)

	var err error
	_handle, err = ConnectPostgreSQL(dsn)
	if err != nil {
		return nil, err
	}
	return _handle, nil
}

// ConnectPostgreSQL
// @docs https://pkg.go.dev/github.com/lib/pq
// dsn e.g.
// user=pqgotest dbname=pqgotest sslmode=verify-full
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
