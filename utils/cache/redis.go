// cache
// redis连接池相关内容

package cache

import (
    "fmt"
    "github.com/alonghub/GinBase/utils/config"
    "github.com/alonghub/GinBase/utils/log"
    "github.com/gomodule/redigo/redis"

    "time"
)

var RedisPool *redis.Pool

// RedisPollInit redis连接处初始化
func RedisPoolInit() {
    RedisPool = &redis.Pool{
        MaxIdle:     config.RedisSetting.MaxIdle,
        MaxActive:   config.RedisSetting.MaxActive,
        IdleTimeout: config.RedisSetting.IdleTimeout,
        Wait:        true,
        Dial: func() (redis.Conn, error) {
            con, err := redis.Dial("tcp", config.RedisSetting.Host,
                //redis.DialPassword(config.RedisSetting["Password"].(string)),
                redis.DialPassword(config.RedisSetting.Password),
                redis.DialDatabase(config.RedisSetting.SelectDB),
                redis.DialConnectTimeout(config.RedisSetting.DialConnectTimeout*time.Second),
                redis.DialReadTimeout(config.RedisSetting.DialReadTimeout*time.Second),
                redis.DialWriteTimeout(config.RedisSetting.DialWriteTimeout*time.Second))
            if err != nil {
                return nil, err
            }
            return con, nil
        },
    }
    _, err := redis.String(RedisPool.Get().Do("ping"))
    if err != nil {
        fmt.Println(err)
        log.Logger.Fatal(err.Error())
    }
}
