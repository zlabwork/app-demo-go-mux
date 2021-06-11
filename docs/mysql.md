
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
