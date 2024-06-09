package db

import "go.mongodb.org/mongo-driver/mongo"


const DBNAME = "youtube-bookmarks"

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}
