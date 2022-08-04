
#### Mysql
```shell
# 查看连接数
SHOW STATUS LIKE 'Threads%';

# 查看最大连接数
SHOW VARIABLES LIKE '%max_connections%';
```

```go
// 缓存连接句柄
var _handle *handle

func getHandle() (*handle, error) {

	if _handle != nil {
		return _handle, nil
	}

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	name := os.Getenv("MYSQL_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", user, pass, host, port, name)

	var err error
	_handle, err = ConnectMySQL(dsn)
	if err != nil {
		return nil, err
	}
	return _handle, nil
}
```


## Postgres
```go
package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type handle struct {
	Conn *sql.DB
}

// ConnectPostgres
// @docs https://pkg.go.dev/github.com/lib/pq
// dsn e.g.
// postgres://username:password@localhost:5432/mydb?sslmode=verify-full
// postgres://username:password@localhost:5432/mydb?sslmode=disable
func ConnectPostgres(dsn string) (*handle, error) {
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

```


## MongoDB
```shell
# 查看连接数
db.serverStatus().connections;
```
