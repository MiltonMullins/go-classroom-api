package repositories

import (
	"context"

	"github.com/miltonmullins/classroom-api/assigment-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StudentTaskRepository interface {
	GetStudentTasks(ctx context.Context, assigmentID int) ([]*models.StudentTask, error)
	CreateStudentTask(ctx context.Context, studentTask *models.StudentTask) error
	UpdateStudentTask(ctx context.Context, studentID, assigmentID int, studentTask *models.StudentTask) error
	DeleteStudentTask(ctx context.Context, studentID, assigmentID int) error
}

type studentTaskRepository struct {
	db *mongo.Client
}

func NewStudentTaskRepository(db *mongo.Client) StudentTaskRepository {
	return &studentTaskRepository{
		db: db,
	}
}

func (r *studentTaskRepository) GetStudentTasks(ctx context.Context, assigmentID int) ([]*models.StudentTask, error) {
	collection := r.db.Database("assigment").Collection("studentTask")
	var studentTask []*models.StudentTask
	filter := bson.D{{Key: "assigment_id", Value: assigmentID}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err //TODO handle error
	}
	if err = cursor.All(ctx, &studentTask); err != nil {
		return nil, err //TODO handle error
	}
	return studentTask, nil

}
func (r *studentTaskRepository) CreateStudentTask(ctx context.Context, studentTask *models.StudentTask) error {
	collection := r.db.Database("assigment").Collection("studentTask")
	_, err := collection.InsertOne(ctx, studentTask)
	if err != nil {
		return err
	}
	return nil
}
func (r *studentTaskRepository) UpdateStudentTask(ctx context.Context, studentID, assigmentID int, studentTask *models.StudentTask) error {
	collection := r.db.Database("assigment").Collection("studentTask")
	filter := bson.D{{Key: "student_id", Value: studentID}, {Key: "assigment_id", Value: assigmentID}}
	update := bson.D{{Key: "$set", Value: studentTask}} //TODO only update Mandatory
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
func (r *studentTaskRepository) DeleteStudentTask(ctx context.Context, studentID, assigmentID int) error {
	collection := r.db.Database("assigment").Collection("studentTask")
	filter := bson.D{{Key: "student_id", Value: studentID}, {Key: "assigment_id", Value: assigmentID}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
