package service

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jeevanshu/warp-url/config"
)

var pool *redis.Pool
var conf *config.Config

// Conn variable to establish connection to redis
var Conn redis.Conn

// URLConfiguration struct to store redis value
type URLConfiguration struct {
	OrignalURL string `redis:"URL"`
	Key        string `redis:"key"`
}

func init() {
	conf = config.New()
	fmt.Println("Connecting to redis")
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.RedisURL, redis.DialDatabase(0))
		},
	}
	Conn = pool.Get()
}
func storeValue(uc URLConfiguration) {

	_, err := Conn.Do(
		"HMSET", "key:"+uc.Key,
		"URL", uc.OrignalURL,
		"key", uc.Key,
	)

	if err != nil {
		log.Printf("Error in executing redis query: %v", err)
	}

}

func fetchValue(minifyURL string) string {
	urlvalues, err := redis.Values(Conn.Do("HGETALL", "key:"+minifyURL))

	if err != nil {
		log.Printf("Error in fetching value from redis %v", err)
	} else if len(urlvalues) == 0 {
		log.Printf("Empty response from redis")
	}
	var url URLConfiguration
	err = redis.ScanStruct(urlvalues, &url)

	if err != nil {
		log.Printf("Error in reading lab %v", err)
	}

	return url.OrignalURL
}
