package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/ramantehlan/redis-ant/internal/config"
	"github.com/ramantehlan/redis-ant/internal/mdb"
	"github.com/ramantehlan/redis-ant/internal/rdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	sourceFile = "ant_source.yml"
)

var ctx = context.Background()

func main() {
	fmt.Println("Client is starting")

	sourceFilePtr := flag.String("source", sourceFile, "Provide the relative path to the source file")
	flag.Parse()

	/////////////////////////////
	// Step 1: Read config file
	/////////////////////////////
	config.InitConfig(*sourceFilePtr)
	fmt.Println("Configs loaded")

	/////////////////////////////
	// Step 2: Connect with mongodb
	/////////////////////////////
	mdb.InitMongoDB()
	fmt.Println("Mongo connected")

	////////////////////////////
	// Step 3: Read mongo
	////////////////////////////
	db := mdb.Client.Database(config.Config.Database)
	coll := db.Collection(config.Config.Collection)

	projection := bson.D{}
	options := options.Find().SetProjection(projection)

	curson, err := coll.Find(context.Background(), bson.D{}, options)
	if err != nil {
		panic(err)
	}

	////////////////////////////
	// Step 4: Write to client cache
	///////////////////////////
	clientRdb := rdb.GetConnection(config.Config.ClientRedisURL, config.Config.ClientRedisPass, 0)
	for curson.Next(ctx) {
		var result bson.M
		err := curson.Decode(&result)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
		keyValue := result[config.Config.KeyField]
		jsonString, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}

		_, redisErr := clientRdb.Do(ctx, "JSON.SET", keyValue, ".", jsonString).Result()
		if redisErr != nil {
			panic(err)
		}
	}

	/////////////////////////////////////////
	// Step 5: Connect with the global redis
	/////////////////////////////////////////
	globalRdb := rdb.GetConnection(config.Config.GlobalRedisURL, config.Config.GlobalRedisPass, 0)
	pubsub := globalRdb.Subscribe(ctx, "ant_updates")

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}
