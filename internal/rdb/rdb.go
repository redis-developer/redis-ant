package rdb

import (
	redis "github.com/go-redis/redis/v8"
)

// GetConnection is
func GetConnection(addr string, password string, database int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       database, // use default DB
	})

	return rdb

}
