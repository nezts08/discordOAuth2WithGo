package repository

import (
	"context"
	"time"

	database "github.com/nezts08/discordOAuth2WithGo/db"
	"github.com/nezts08/discordOAuth2WithGo/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func collection() *mongo.Collection {
	return database.DB.Collection("users")
}

func CreateUser(user *models.UserDC) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return collection().InsertOne(ctx, user)
}

func FindUserByDiscordID(discordID string) (*models.UserDC, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.UserDC
	err := collection().
		FindOne(ctx, bson.M{"id": discordID}).
		Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByID(id string) (*models.UserDC, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.UserDC
	err = collection().
		FindOne(ctx, bson.M{"_id": objID}).
		Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindAllUsers() ([]models.UserDC, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.UserDC
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(discordID string, update bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return collection().UpdateOne(
		ctx,
		bson.M{"id": discordID},
		bson.M{"$set": update},
	)
}

func UpdateUserFull(user *models.UserDC) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return collection().UpdateOne(
		ctx,
		bson.M{"id": user.ID},
		bson.M{"$set": user},
	)
}

func DeleteUserByDiscordID(discordID string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return collection().DeleteOne(ctx, bson.M{
		"id": discordID,
	})
}

func DeleteUserByID(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return collection().DeleteOne(ctx, bson.M{
		"_id": objID,
	})
}
