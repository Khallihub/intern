package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main() {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	err = mongoClient.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	defer mongoClient.Disconnect(context.Background())

	server := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}

	server.Use(cors.New(config))

	server.Use(gin.Recovery(), gin.Logger())

	controller := NewTaskController(mongoClient, "testDb", "todos")

	server.POST("/create", controller.CreateTask)
	server.GET("/getall", controller.GetAllTasks)
	server.POST("/getone", controller.GetTaskById)
	server.PUT("/update", controller.UpdateTask)
	server.Delete("/delete", controller.DeleteTask)

}