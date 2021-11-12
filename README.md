## 最佳实践
按依赖划分项目结构


## request
```go
// 获取 uri
vars := mux.Vars(r)
vars["name"]

// 获取参数 GET
vars := r.URL.Query()
vars.Get("name")

// 获取参数 POST PUT PATCH 
// application/x-www-form-urlencoded
name := r.PostFormValue("name")

// 获取参数 GET POST PUT PATCH
name := r.FormValue("name")

// 获取 body 任意请求类型
var reader io.Reader = r.Body
b, _ := ioutil.ReadAll(reader)
v, _ := url.ParseQuery(string(b))
```


## cookie
```go
// 读取 cookie
r.Cookies()

// 写入 cookie
cookie := &http.Cookie{Name: "userId", Value: "123456"}
http.SetCookie(w, cookie)
```


## 优秀组件
| 包名 | 简介 |
| --- | --- |
| [github.com/gorilla/mux](https://github.com/gorilla/mux) | web 框架 |
| [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) | web 框架 |
| [github.com/gorilla/sessions](https://github.com/gorilla/sessions) | session |
| [github.com/joho/godotenv](https://github.com/joho/godotenv) | 配置, env 环境变量 |
| [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) | 配置, yaml |
| [github.com/spf13/viper](https://github.com/spf13/viper) | 配置, 支持多种格式 |
| [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) | 数据库驱动, MySQL |
| [github.com/gocql/gocql](https://github.com/gocql/gocql) | 数据库驱动, Cassandra |
| [go.mongodb.org/mongo-driver](https://go.mongodb.org/mongo-driver) | 数据库驱动, mongoDB |
| [github.com/go-redis/redis](https://github.com/go-redis/redis) | 数据库驱动, redis |
| [github.com/gomodule/redigo](https://github.com/gomodule/redigo) | 数据库驱动, redis 不推荐 |
| [github.com/go-resty/resty/](https://github.com/go-resty/resty/) | http 请求客户端 |
| [github.com/valyala/fastjson](https://github.com/valyala/fastjson) | json 解析 |
| [github.com/tidwall/gjson](https://github.com/tidwall/gjson) | json 解析, 注意: []强制转为字符串后为[] |
| [github.com/buger/jsonparser](https://github.com/buger/jsonparser) | json 解析 |
| [github.com/Sirupsen/logrus](https://github.com/Sirupsen/logrus) | 日志 |
| [github.com/uber-go/zap](https://github.com/uber-go/zap) | 日志 |
| [github.com/bwmarrin/snowflake](https://github.com/bwmarrin/snowflake) | ID生成 |
| [github.com/google/uuid](https://github.com/google/uuid) | UUID生成 |
| [github.com/silenceper/pool](https://github.com/silenceper/pool) | 线程池 |


## 技术栈
| 技术 | 简介 |
| --- | --- |
| [Zipkin](https://zipkin.io/) | 链路追踪 |
| [Jaeger](https://www.jaegertracing.io/)   | 链路追踪 |


## 工具
https://mholt.github.io/json-to-go/


## 参考文档
https://github.com/avelino/awesome-go  
https://zhuanlan.zhihu.com/p/39326315  
https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1  
https://studygolang.com/articles/17467?fr=sidebar 

