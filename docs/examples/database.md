
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
```sql
-- 查看当前配置的最大连接数
show max_connections;

-- 查询当前活跃连接数
select count(1) from pg_stat_activity;
```

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


## ORM - gorm
```shell
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

```go
package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

// Connect
// @docs https://github.com/go-sql-driver/mysql#dsn-data-source-name
// @docs https://gorm.io/zh_CN/docs/connecting_to_the_database.html
// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func Connect(dsn string) (*gorm.DB, error) {

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := conn.DB()

	// 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)

	// 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)

	return conn, err
}

func getHandle() (*gorm.DB, error) {

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	name := os.Getenv("MYSQL_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local", user, pass, host, port, name)
	return Connect(dsn)
}
```
