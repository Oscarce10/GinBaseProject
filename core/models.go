package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Priority    int                `json:"priority"`
	DueDate     string             `json:"due_date"`
}

func getTask(task *Task) map[string]interface{} {
	return map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"priority":    task.Priority,
		"due_date":    task.DueDate,
	}
}
