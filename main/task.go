package main

type Task struct {
	ID     string        `json:"id" bson:"_id,omitempty"`
	Title string        `json:"title" binding:"required"`
	Description string      `json:"description" bson:"description"`
	DueDate string      `json:"dueDate" bson:"dueDate"`
}
