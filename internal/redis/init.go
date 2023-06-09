package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"hzer/configs"
	"time"
)

var pool *redis.Pool

func InitRedis(cf configs.Redis) {
	pool = &redis.Pool{
		// Maximum number of connections allocated by the pool at a given time.
		// When zero, there is no limit on the number of connections in the pool.
		//最大活跃连接数，0代表无限
		MaxActive: 1000,
		//最大闲置连接数
		// Maximum number of idle connections in the pool.
		MaxIdle: 50,
		//闲置连接的超时时间
		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: time.Second * 100,
		//定义拨号获得连接的函数
		// Dial is an application supplied function for creating and configuring a
		// connection.
		//
		// The connection returned from Dial must not be in a special state
		// (subscribed to pubsub channel, transaction started, ...).
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cf.Host, cf.Port))
			if err != nil {
				return nil, err
			}
			if cf.Password != "" {
				if _, err := c.Do("AUTH", cf.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}
