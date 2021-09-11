package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init()  {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			setPasswd := redis.DialPassword("123456")
			return redis.Dial("tcp", "localhost:6379", setPasswd)
		},
		MaxIdle:         8,
		MaxActive:       0,
		IdleTimeout:     100,
	}
}
func main() {

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "wang")
	if err !=nil{
		fmt.Println(err)
		return
	}

	reply, err := redis.String(conn.Do("get", "name"))
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

}