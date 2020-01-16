package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	_mongoURL      = os.Getenv("MONGO_URL") // mongodb://myDBReader:D1fficultP%40ssw0rd@mongodb0.example.com:27017/admin
	_mongoDatabase = os.Getenv("MONGO_DATABASE")
)

type IMongoDriver interface {
	DB() *mongo.Database
}

type Driver struct {
	db *mongo.Database
}

func NewMongoDriver() IMongoDriver {
	result := &Driver{}
	if err := result.establishConnection(); err != nil {
		log.Fatal(err)
	}
	return result
}

func (d *Driver) DB() *mongo.Database {
	return d.db
}

func (d *Driver) establishConnection() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(_mongoURL))
	if err != nil {
		return fmt.Errorf("driver.EstablishConnection() on NewClient:%w", err)
	}
	if err := client.Connect(ctx); err != nil {
		return fmt.Errorf("driver.EstablishConnection() on Connect:%w", err)
	}
	d.db = client.Database(_mongoDatabase)
	return nil
}
