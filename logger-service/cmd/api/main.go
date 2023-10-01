package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPost  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPost = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {

	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create a contex in order to disconnect
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// start server
	log.Println("Server started at port:", webPort)
	// go app.serve()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

// func (app *Config) serve() {
// 	srv := &http.Server{
// 		Addr:    fmt.Sprintf(":%s", webPort),
// 		Handler: app.routes(),
// 	}

// 	err := srv.ListenAndServe()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connection:", err)
		return nil, err
	}

	log.Println("Connected to mongoDB!")

	return c, nil
}
