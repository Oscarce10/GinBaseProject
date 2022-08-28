package core

import (
	config "GinBaseProject/config/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	DueDate     string `json:"due_date"`
}

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	DueDate     string `json:"due_date"`
}

func FindTasks() []interface{} {
	var db config.Database
	db.Client = db.GetClient()
	collection := db.GetCollection("tasks")
	defer func() {
		db.CloseClient()
	}()
	var tasks []interface{}
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal("Error finding tasks: ", err.Error())
	}
	for cursor.Next(context.Background()) {
		var task Task
		cursor.Decode(&task)
		tasks = append(tasks, getTask(&task))
	}
	return tasks
}

func CreateTask(title, description, dueDate string, priority int) {
	var db config.Database
	db.Client = db.GetClient()
	collection := db.GetCollection("tasks")
	defer func() {
		db.CloseClient()
	}()
	task := Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		DueDate:     dueDate,
	}
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal("Error creating task: ", err.Error())
	}
	return
}
