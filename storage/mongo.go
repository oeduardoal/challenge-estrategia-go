package storage

import (
	"challenge-go-react/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// DB define struct
type DB struct{}

// Connect define a var to connect
func (m *DB) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	log.Print("connected!")
}

// FindUserByUsername response
func (m *DB) FindUserByUsername(username string) (models.User, error) {
	collection := client.Database("golang").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var user models.User
	collection.FindOne(ctx, bson.M{"username": bson.M{"$eq": username}}).Decode(&user)
	if user.Username != "" {
		return user, nil
	} else {
		return models.User{}, errors.New("not found")
	}
}

// InsertUser reponse
func (m *DB) InsertUser(newUser *models.User) *mongo.InsertOneResult {
	collection := client.Database("golang").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, _ := collection.InsertOne(ctx, newUser)
	return result
}

// FindTagsByUsernameAndReponame response
func (m *DB) FindTagsByUsernameAndReponame(username string, reponame string) (models.TagRequest, error) {
	collection := client.Database("golang").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var user models.User
	collection.FindOne(ctx, bson.M{"username": bson.M{"$eq": username}, "repos.name": bson.M{"$eq": reponame}}).Decode(&user)
	if user.Username != "" {
		var tags models.TagRequest
		for index := 0; index < len(user.Repos); index++ {
			if user.Repos[index].Name == reponame {
				tags = models.TagRequest{Reponame: user.Repos[index].Name, Tags: user.Repos[index].Tags}
				break
			}
		}
		return tags, nil
	} else {
		return models.TagRequest{}, errors.New("not found")
	}
}

// FindTagsAndUpdate response
func (m *DB) FindTagsAndUpdate(username string, reponame string, tags models.TagRequest) (models.TagRequest, error) {
	collection := client.Database("golang").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var user models.User
	myoptions := options.FindOneAndUpdateOptions{}
	myoptions.SetReturnDocument(options.After)
	collection.FindOneAndUpdate(ctx, bson.M{"username": username, "repos.name": reponame}, bson.M{"$set": bson.M{"repos.$.tags": tags.Tags}}, &myoptions).Decode(&user)
	if user.Username != "" {
		var tags models.TagRequest
		for index := 0; index < len(user.Repos); index++ {
			if user.Repos[index].Name == reponame {
				tags = models.TagRequest{Reponame: user.Repos[index].Name, Tags: user.Repos[index].Tags}
				break
			}
		}
		return tags, nil
	} else {
		return models.TagRequest{}, errors.New("not found")
	}
}

// InsertRec reponse
func (m *DB) InsertRec(recList *models.RecListRequest) *mongo.InsertOneResult {
	collection := client.Database("golang").Collection("rec")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, _ := collection.InsertOne(ctx, recList)
	return result
}

// FindRecByQuery response
func (m *DB) FindRecByQuery(search string) (models.RecListRequest, error) {
	collection := client.Database("golang").Collection("rec")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var rec models.RecListRequest
	filter := bson.M{"$or": bson.A{bson.M{"search": search}, bson.M{"rec.word": bson.M{"$regex": search}}}}
	collection.FindOne(ctx, filter).Decode(&rec)
	if rec.Search != "" {
		rec.Search = search
		return rec, nil
	} else {
		return models.RecListRequest{}, errors.New("not found")
	}
}

// FindRepoByQuery response
func (m *DB) FindRepoByQuery(username string, search string) ([]*models.Repo, error) {
	collection := client.Database("golang").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// filter := bson.A{bson.M{"$unwind": "$repos"}, bson.M{"$match": bson.M{"repos.tags": bson.M{"$in": bson.A{search}}}}}
	filter := bson.A{bson.M{"$unwind": "$repos"}, bson.M{"$match": bson.M{"$and": bson.A{bson.M{"repos.tags": bson.M{"$in": bson.A{"top"}}}, bson.M{"username": username}}}}}
	cur, _ := collection.Aggregate(ctx, filter)
	var results []*models.Repo
	for cur.Next(context.TODO()) {
		var newUser struct {
			ID   interface{} `json:"id,omitempty" bson:"_id,omitempty"`
			Repo models.Repo `json:"repos,omitempty" bson:"repos,omitempty"`
		}
		err := cur.Decode(&newUser)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &newUser.Repo)
	}

	return results, nil

}
