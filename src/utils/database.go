package utils

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	instance *mongo.Client
)

// Db
func Db() *mongo.Database {
  conf := Config()
	once.Do(func() {
		var err error

		instance, err = initClient(conf.Dsn, 10)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance.Database(conf.DBName)
}

func initClient(dsn string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseClient(timeout int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	var err error
	if instance != nil {
		err = instance.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("instace is nil database instance closing before initialization")
	}
}
