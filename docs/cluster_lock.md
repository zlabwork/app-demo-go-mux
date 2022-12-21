## 分布式锁

```go
Lock(r.Context(), client, "lock_key_name", func() {
    fmt.Println(abc)
})
```

```go
func Lock(ctx context.Context, cli *redis.Client, lockKey string, something func()) error {

    val := rand.Intn(1 << 31)
    exp := 30 * time.Second

    i := 0
    for {
        isLock := cli.SetNX(ctx, lockKey, val, exp).Err() == nil
        if isLock {
            v, _ := cli.Get(ctx, lockKey).Int()
            if v == val {
                something()
                cli.Del(ctx, lockKey)
                return nil
            }
        } else {
            time.Sleep(500 * time.Millisecond)
        }

        i++
        if i > 100 {
            break
        }
    }

    return fmt.Errorf("redis lock is failed")
}
```

[redis加锁的几种实现](http://ukagaka.github.io/php/2017/09/21/redisLock.html)
[怎样实现redis分布式锁](https://www.zhihu.com/question/300767410?sort=created)
[分布式锁之Redis实现](https://www.jianshu.com/p/47fd7f86c848)
