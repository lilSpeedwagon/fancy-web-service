package database

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	databaseName   = "defaultDB"
	collectionName = "entries"

	mongoKey    = "key"
	mongoValue  = "value"
	mongoCmdSet = "$set"

	shortTimeout   = 1 * time.Second
	requestTimeout = 5 * time.Second
)

type dbMongo struct {
	client     *mongo.Client
	collection *mongo.Collection
	context    context.Context
}

func (db dbMongo) timeoutContext(t time.Duration) context.Context {
	ctx, _ := context.WithDeadline(db.context, time.Now().Add(t))
	return ctx
}

func (db dbMongo) Put(key string, value string) (bool, error) {
	filter := bson.D{{mongoKey, key}}
	updFilter := bson.D{
		{mongoCmdSet, bson.D{
			{mongoValue, value},
		}},
	}

	updResult, updErr := db.collection.UpdateOne(db.timeoutContext(shortTimeout), filter, updFilter)
	if updErr != nil {
		return false, updErr
	}

	// key already exists
	if updResult.MatchedCount != 0 {
		return false, nil
	}

	data := bson.D{
		{mongoKey, key},
		{mongoValue, value},
	}

	_, insErr := db.collection.InsertOne(db.timeoutContext(shortTimeout), data)
	if insErr != nil {
		return false, insErr
	}

	return true, nil
}

func (db dbMongo) Remove(key string) (bool, error) {
	filter := bson.D{{mongoKey, key}}

	delRes, err := db.collection.DeleteOne(db.timeoutContext(shortTimeout), filter)
	if err != nil {
		return false, nil
	}

	return delRes.DeletedCount != 0, nil
}

func (db dbMongo) Read(key string) (string, error) {
	filter := bson.D{{mongoKey, key}}

	findRes := db.collection.FindOne(db.timeoutContext(shortTimeout), filter)
	if findRes.Err() == mongo.ErrNoDocuments {
		return "", nil
	}

	var data bson.D
	if err := findRes.Decode(&data); err != nil {
		return "", err
	}

	value := fmt.Sprintf("%v", data.Map()[mongoValue])
	return value, nil
}

func (db dbMongo) Close() error {
	return db.client.Disconnect(db.timeoutContext(requestTimeout))
}

func makeMongoDB(url string) (IDataBase, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	mainContext := context.Background()
	ctx, _ := context.WithDeadline(mainContext, time.Now().Add(requestTimeout))

	//noinspection ALL
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf(
			"cannot find collection %s in database %s", collectionName, databaseName))
	}

	return dbMongo{client, collection, mainContext}, nil
}
