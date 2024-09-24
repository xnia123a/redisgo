package cache

import (
	"github.com/garyburd/redigo/redis"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	Pool *redis.Pool
)

const (
	Host = ""
	Port = 6379
	User = "root"
	Password = ""
	Db = 4
)

func init() {
	redisHost := Host + ":" + strconv.Itoa(Port)
	Pool = newPool(redisHost)
	cleanupHook()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialPassword(Password), redis.DialDatabase(Db))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
