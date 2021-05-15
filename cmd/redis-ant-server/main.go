package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"time"

	"github.com/rwynn/gtm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ramantehlan/redis-ant/internal/config"
	"github.com/ramantehlan/redis-ant/internal/rdb"
)

const (
	sourceFile = "ant_source.yml"
)

var redisCtx = context.Background()

func main() {
	fmt.Println("Server is starting")

	sourceFilePtr := flag.String("source", sourceFile, "Provide the relative path to the source file")
	flag.Parse()

	/////////////////////////////
	// Step 1: Read config file
	/////////////////////////////
	config.InitConfig(*sourceFilePtr)
	fmt.Println("Configs loaded")

	///////////////////////////////
	// Step 2: We want to listen to the OpLogs from MongoDB
	///////////////////////////////
	rb := bson.NewRegistryBuilder()
	rb.RegisterTypeMapEntry(bsontype.DateTime, reflect.TypeOf(time.Time{}))
	reg := rb.Build()
	clientOptions := options.Client()
	clientOptions.SetRegistry(reg)
	mongoSTR := "mongodb+srv://" + config.Config.MongoUser + ":" + config.Config.MongoPass + "@" + config.Config.MongoURI
	clientOptions.ApplyURI(mongoSTR)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	ctxm, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctxm)
	if err != nil {
		panic(err)
	}

	collection := config.Config.Database + "." + config.Config.Collection
	defer client.Disconnect(context.Background())
	ctx := gtm.Start(client, &gtm.Options{
		//DirectReadNs:   []string{collection},
		ChangeStreamNs: []string{collection},
		//MaxWaitSecs:    10,
		OpLogDisabled: true,
	})

	////////////////////////////////////////////
	// Step 3: Publish change to redis
	//////////////////////////////////////////
	globalRdb := rdb.GetConnection(config.Config.GlobalRedisURL, config.Config.GlobalRedisPass, 0)
	for {
		select {
		case err := <-ctx.ErrC:
			fmt.Printf("got err %+v", err)
			break
		case op := <-ctx.OpC:

			fmt.Printf("got op %+v", op.Data)
			message, jErr := json.Marshal(op.Data)
			if jErr != nil {
				panic(err)
			}

			err := globalRdb.Publish(redisCtx, "ant_updates", message).Err()
			if err != nil {
				panic(err)
			}
			break
		}
	}
}
