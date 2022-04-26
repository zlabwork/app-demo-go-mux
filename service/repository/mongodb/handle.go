package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type handle struct {
	Conn *mongo.Client
}

// ConnectMongodb
// https://docs.mongodb.com/drivers/go/current/
// https://developer.mongodb.com/community/forums/tag/golang/
// https://www.mongodb.com/docs/drivers/go/current/fundamentals/connection/
//
// URI e.g.
// mongodb://127.0.0.1:27017
// mongodb://user:pass@hostname:27017
// mongodb://user:pass@hostname:27017/?maxPoolSize=20&w=majority
// mongodb://host1:27017,host2:27017,host3:27017/?replicaSet=myRS
func ConnectMongodb(uri string) (*handle, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Client().ApplyURI(uri)
	cli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	err = cli.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &handle{
		Conn: cli,
	}, nil
}
