package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	domain "resource_service/internal/core/domain/resource"
	"resource_service/internal/core/helper"
	port "resource_service/internal/ports/resource"
	"time"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (port.ResourceRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}
func (r *mongoRepository) Read(reference string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	resource := domain.Resource{} // i removed & from here
	collection := r.client.Database(r.database).Collection("resources")
	filter := bson.M{"reference": reference}
	err := collection.FindOne(ctx, filter).Decode(&resource)
	if err != nil {
		return nil, helper.PrintErrorMessage("404", err.Error())
	}
	return resource, nil
}
func (r *mongoRepository) ReadAll() (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	var resources []domain.Resource
	var resource domain.Resource
	collection := r.client.Database(r.database).Collection("resources")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, helper.PrintErrorMessage("404", err.Error())
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&resource)
		if err != nil {

			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}
func (r *mongoRepository) Create(resource domain.Resource) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("resources")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"reference":     resource.Reference,
			"created_on":    resource.CreatedOn,
			"last_modified": resource.LastModified,
			"name":          resource.Name,
			"value":         resource.Value,
		},
	)
	if err != nil {

		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	return resource.Reference, nil
}
func (r *mongoRepository) Update(reference string, resource domain.Resource) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("resources")
	_, err := collection.ReplaceOne(
		ctx,
		bson.M{"reference": reference},
		bson.M{
			"reference":     resource.Reference,
			"created_on":    resource.CreatedOn,
			"last_modified": resource.LastModified,
			"name":          resource.Name,
			"value":         resource.Value,
		},
	)
	if err != nil {
		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	return resource.Reference, nil
}
func (r *mongoRepository) Delete(reference string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("resources")
	_, err := collection.DeleteOne(
		ctx,
		bson.M{"reference": reference},
	)
	if err != nil {

		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	return reference, nil
}
