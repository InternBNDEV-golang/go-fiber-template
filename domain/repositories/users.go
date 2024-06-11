package repositories

import (
	"context"
	. "go-fiber-template/domain/datasources"
	"go-fiber-template/domain/entities"
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type usersRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IUsersRepository interface {
	InsertUser(data entities.UserDataModel) error
	FindAll() (*[]entities.UserDataModel, error)
}

func NewUsersRepository(db *MongoDB) IUsersRepository {
	return &usersRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("users"),
	}
}

func (repo *usersRepository) InsertUser(data entities.UserDataModel) error {
	if _, err := repo.Collection.InsertOne(repo.Context, data); err != nil {
		fiberlog.Errorf("Users -> InsertNewUser: %s \n", err)
		return err
	}
	return nil
}

func (repo *usersRepository) FindAll() (*[]entities.UserDataModel, error) {
	options := options.Find()
	filter := bson.M{}
	var users []entities.UserDataModel

	cursor, err := repo.Collection.Find(repo.Context, filter, options)
	if err != nil {
		fiberlog.Errorf("Users -> FindAll: %s \n", err)
		return nil, err
	}
	defer cursor.Close(repo.Context)

	err = cursor.All(repo.Context, &users)
	if err != nil {
		fiberlog.Errorf("Users -> FindAll: %s \n", err)
		return nil, err
	}

	return &users, nil
}
