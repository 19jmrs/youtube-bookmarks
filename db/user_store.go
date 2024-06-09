package db

import (
	"context"

	"github.com/19jmrs/youtube-bookmarks/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}




func NewMongoUserStore(client *mongo.Client) *MongoUserStore{

	return &MongoUserStore{
		client: client,
		coll: client.Database(DBNAME).Collection(userColl),
	}
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error){
	res, err := s.coll.InsertOne(ctx, user)
	if err!= nil{
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}