package repositories

import (
	"context"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	//"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AssigmentRepository interface {
	GetAssigment(ctx context.Context, param string) ([]*models.Assigment, error)
	CreateAssigment(ctx context.Context, assigment *models.Assigment) error
	UpdateAssigment(ctx context.Context, assigmentID int, assigment *models.Assigment) error
	DeleteAssigment(ctx context.Context, assigmentID int) error
}

type assigmentRepository struct {
	db *mongo.Client
}

func NewAssigmentRepository(db *mongo.Client) AssigmentRepository {
	return &assigmentRepository{
		db: db,
	}
}

func (r *assigmentRepository) GetAssigment(ctx context.Context, param string) ([]*models.Assigment, error) {
	collection := r.db.Database("assigment").Collection("assigment")
	var assigments []*models.Assigment
	filter := bson.D{{Key: "title", Value: param}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err //TODO handle error
	}
	if err = cursor.All(ctx, &assigments); err != nil {
		return nil, err //TODO handle error
	}
	return assigments, nil
}

func (r *assigmentRepository) CreateAssigment(ctx context.Context, assigment *models.Assigment) error {
	collection := r.db.Database("assigment").Collection("assigment")
	_, err := collection.InsertOne(ctx, assigment)
	if err != nil {
		return err
	}
	return nil
}

func (r *assigmentRepository) UpdateAssigment(ctx context.Context, assigmentID int, assigment *models.Assigment) error {
	collection := r.db.Database("assigment").Collection("assigment")
	filter := bson.D{{Key: "ID", Value: assigmentID}}
	update := bson.D{{Key: "$set", Value: assigment}}
	/*
	   	upsert := true
	   	after := options.After
	    	opt := options.FindOneAndUpdateOptions{
	   		ReturnDocument: &after,
	   		Upsert:         &upsert,
	   		}
	   		result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	*/

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *assigmentRepository) DeleteAssigment(ctx context.Context, assigmentID int) error {
	collection := r.db.Database("assigment").Collection("assigment")
	filter := bson.D{{Key: "ID", Value: assigmentID}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
