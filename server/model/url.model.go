package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Url struct {
	LUrl         string    `json:"lUrl,omitempty" bson:"lUrl"`
	SUrl         string    `json:"sUrl"  bson:"sUrl"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	LastViewed   time.Time `json:"lastViewed" bson:"lastViewed"`
	CountVisited int       `json:"countVisited" bson:"countVisited"`
	UserIp       string    `json:"userIp" bson:"userIp"`
}

func AddUrlToDb(u Url) (string, error) {
	col := database.Db.Collection("shorturls")

	i, err := col.InsertOne(database.Ctx, u)
	if err != nil {
		return "", errors.New("insert Failed")
	}

	str, err := col.Indexes().CreateOne(database.Ctx, mongo.IndexModel{Keys: bson.D{{Key: "sUrl", Value: 1}}, Options: options.Index().SetUnique(true)})
	fmt.Printf("\n  adding indexing %+v or error %+v\n", str, err)

	return i.InsertedID.(primitive.ObjectID).Hex(), nil
}

func FindUrlBySUrl(sUrl string) (Url, error) {
	col := database.Db.Collection("shorturls")

	var url Url
	filter := bson.D{{Key: "sUrl", Value: sUrl}}

	if err := col.FindOne(database.Ctx, filter).Decode(&url); err != nil {
		return url, errors.New("no URL found")
	}
	return url, nil
}
