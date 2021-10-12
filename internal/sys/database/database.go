package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Set of error variables for CRUD operations.
var (
	ErrDBNotFound = errors.New("not found")
)

type Config struct {
	URI          string
	DatabaseName string
	Ctx          context.Context
}

func Open(cfg Config) (*mongo.Client, error) {
	clientOptions := options.Client().
		ApplyURI(cfg.URI)
	ctx, cancel := context.WithTimeout(cfg.Ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	// defer client.Disconnect(ctx)
	return client, nil
	// database := client.Database(cfg.DatabaseName)
	// return database, nil
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, client *mongo.Client) error {

	// First check we can ping the database.
	var pingError error
	for attempts := 1; ; attempts++ {
		pingError = client.Ping(ctx, readpref.Primary())
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	// Make sure we didn't timeout or be cancelled.
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// TODO: Determine the MongoDB equivalent of this
	// Run a simple query to determine connectivity. Running this query forces a
	// round trip through the database.
	// const q = `SELECT true`
	// var tmp bool
	// return sqlxDB.QueryRowContext(ctx, q).Scan(&tmp)
	return nil
}
