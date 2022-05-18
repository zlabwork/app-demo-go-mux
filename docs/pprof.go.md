## Http
#### pprof - http server [ default http server ]
```go
import (
    _ "net/http/pprof"
)
```

#### pprof - http server [ gorilla/mux ]

```go
// import "net/http/pprof"
// 自定义的 Mux，则需要手动注册一些路由规则
mux.HandleFunc("/debug/pprof/", pprof.Index)
mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

## Not http
#### cpu
```go
import "runtime/pprof"
```

```go
f, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
defer f.Close()
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
// do something
```

#### memory
```go
f, _ := os.OpenFile("mem.profile", os.O_CREATE|os.O_RDWR, 0644)
defer f.Close()
// do something
pprof.Lookup("heap").WriteTo(f, 0)
```


## Profiling
```shell
# 分析文件
go tool pprof cpu.profile
go tool pprof mem.profile

top
list <func>
```


```shell
# e.g.
root@zlab:~# go tool pprof mem.profile
Type: inuse_space
Time: May 18, 2022 at 2:31pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1544.52kB, 100% of 1544.52kB total
Showing top 10 nodes out of 11
      flat  flat%   sum%        cum   cum%
 1032.02kB 66.82% 66.82%  1032.02kB 66.82%  main.repeat (inline)
  512.50kB 33.18%   100%   512.50kB 33.18%  runtime.allocm
         0     0%   100%  1032.02kB 66.82%  main.main
         0     0%   100%  1032.02kB 66.82%  runtime.main
         0     0%   100%   512.50kB 33.18%  runtime.mcall
         0     0%   100%   512.50kB 33.18%  runtime.newm
         0     0%   100%   512.50kB 33.18%  runtime.park_m
```

## Docs
https://segmentfault.com/a/1190000040152398  
https://www.jianshu.com/p/6175798c03b4  
