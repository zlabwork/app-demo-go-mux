package mongodb

import (
    "app/configs"
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

type handle struct {
    Conn     *mongo.Client
    Database *mongo.Database
}

// mongodb://127.0.0.1:27017
// mongodb://foo:bar@localhost:27017
func ConnectMongodb(c *configs.MongoConfig) (*handle, error) {
    dsn := fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    opts := options.Client().ApplyURI(dsn)
    cli, err := mongo.Connect(ctx, opts)
    if err != nil {
        return nil, err
    }
    err = cli.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }

    return &handle{
        Conn:     cli,
        Database: cli.Database(c.Name),
    }, nil
}
