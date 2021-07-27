package users

import (
	"context"
	"flatApp/pkg/platform/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, u user.User) (interface{}, error) {
	database := r.db.Database("users_db")

	collection := database.Collection("users")

	res, err := collection.InsertOne(ctx, u)

	if err != nil {
		return user.User{}, err
	}

	id := res.InsertedID
	return id, nil
}

func (r *Repository) GetUser(ctx context.Context, username, password string) (user.User, error) {
	var usr user.User

	collection := r.db.Database("users_db").Collection("users")

	err := collection.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&usr)
	if err != nil {
		return usr, err
	}

	return usr, nil
}