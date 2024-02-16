package main

import (
	"context"
	"fmt"
	// "errors"
	// "fmt"
	"github.com/khallihub/godoc/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService interface {
	GetAllTasks() ([]Task, error)
	GetTask(id string) (Task, error)
	CreateNewTask(ctx *gin.Context)
	UpdateTask(taskId string, body dto.DocumentData) error
	DeleteTask(ctx *gin.Context) (*dto.Document, error)
}

type taskService struct {
	collection *mongo.Collection // MongoDB collection
}

func NewTaskService(client *mongo.Client, databaseName, collectionName string) TaskService {
	collection := client.Database(databaseName).Collection(collectionName)
	return &taskService{
		collection: collection,
	}
}

func (service *documentService) GetAllTasks() ([]*task, error) {
	cursor, err := service.collection.Find(context.Background())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var tasks []*Task
	for cursor.Next(context.Background()) {
		var document dto.Document
		if err := cursor.Decode(&tasks); err != nil {
			return nil, err
		}
		taskDto := &Task{
			ID:    task.ID,
			Title: task.Title,
			Description: task.description,
			DueDate: task.dueDate
		}

		tasks = append(documents, documentDTO)
	}
	return tasks, nil
}

func (service *documentService) GetTask(id string) ([]*Task, error) {
	// Implement logic to fetch a document by its ID from the MongoDB collection
	var taskID string
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	err = service.collection.FindOne(context.Background(), filter).Decode(&taskID)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

