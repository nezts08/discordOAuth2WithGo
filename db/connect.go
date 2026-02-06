package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverApi := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().
		ApplyURI("mongodb+srv://" +
			os.Getenv("DB_USER") + ":" +
			os.Getenv("DB_PASSWORD") +
			"@cluster0.tuwdk6f.mongodb.net/?appName=Cluster0").
		SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	Client = client
	DB = client.Database("Testing")

	fmt.Println("✅ MongoDB conectado")
}
