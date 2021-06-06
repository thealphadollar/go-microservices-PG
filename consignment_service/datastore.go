package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	conn, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err := conn.Ping(ctx, nil); err != nil {
		if retry == 0 {
			return nil, err
		}
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry-1)
	}
	return conn, nil
}