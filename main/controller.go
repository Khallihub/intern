package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DocumentController interface {
	GetAllTasks(ctx *gin.Context) ([]Task, error)
	GetTask(ctx *gin.Context) (Task, error)
	CreateNewTask(ctx *gin.Context)
	UpdateTask(taskId string, body dto.DocumentData) error
	DeleteTask(ctx *gin.Context) (*dto.Document, error)
}

type taskController struct {
	taskService  TaskService
}

func NewTaskController(taskService TaskService) TaskController {
	return &taskController{
		taskService: taskService,
	}
}

func (controller *taskController) GetAllTasks(ctx *gin.Context) ([]*Task, error) {
	// Implement logic to fetch all documents from the MongoDB collection of a single user
	tasks, err := taskService.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return nil, err
	}

	return tasks, nil
}

func (contrller *taskController) GetTaskById(ctx *gin.Context) (*Task, error) {
	var id string
	err := ctx.ShouldBind(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return nil, err
	}
	task, err := GetTaskById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search tasks"})
		return nil, err
	}

	return task, nil
}

func (controller *taskController) CreateNewTask(ctx *gin.Context) {
	var task Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	
	taskId, err := taskService.CreateDocument(task.Title, task.Description, task.DueDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "task created successfully", "task_id": taskId})
}

func (controller *taskController) UpdateTask(documentID string, body Task) error {
	err := taskService.UpdateDocument(documentID, body)
	if err != nil {
		return err
	}
	return nil
}

func (controller *taskController) DeleteTask(ctx *gin.Context) {
	taskID := ctx.Param("id")	

	if taskID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	s, err := taskService.Deletetask(taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	if s == false {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}