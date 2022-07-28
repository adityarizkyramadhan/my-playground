package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(client.Get("key").Val())

	/*
		result when I don't start the redis server
		dial tcp [::1]:6379: connectex: No connection could be made because the target machine actively refused it.
		panic: dial tcp [::1]:6379: connectex: No connection could be made because the target machine actively refused it.
	*/
	/*
		-how to start the redis server
		sudo service redis-server start

		-check the redis-server status
		sudo service redis-server status

		-stop the redis-server
		sudo service redis-server stop

		-if want to use redis-server in the background
		sudo service redis-server start --daemonize yes

		-if want to use redis-cli to interact with the redis-server
		redis-cli



	*/
}
