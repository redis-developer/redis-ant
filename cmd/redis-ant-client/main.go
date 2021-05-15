package main

import (
	"context"
	"flag"
	"fmt"
	redis "github.com/go-redis/redis/v8"

	"github.com/ramantehlan/redis-ant/internal/config"
)

const (
	sourceFile = "ant_source.yml"
)

var ctx = context.Background()

func main() {
	fmt.Println("Hello from Client")

	sourceFilePtr := flag.String("source", sourceFile, "Provide the relative path to the source file")
	flag.Parse()

	// Step 1: Read config file
	config.InitConfig(*sourceFilePtr)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

}
