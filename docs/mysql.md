
#### 查询1行
```go
dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASS"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_NAME"),
)
db, _ := sql.ConnectMySQL(dsn)
var id int64
var path string
row := db.Conn.QueryRow("SELECT id, path FROM `files` LIMIT 1")
row.Scan(&id, &path)
fmt.Println(id, path)
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
