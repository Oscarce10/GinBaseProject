package core

import (
	"log"
	"net/http"
)

func GetTasksProcess() map[string]interface{} {
	tasks := FindTasks()
	if len(tasks) > 0 {
		log.Println("Found tasks: ", tasks)
		return map[string]interface{}{
			"status_code": http.StatusOK,
			"data":        tasks,
		}
	} else {
		log.Println("No tasks found")
		return map[string]interface{}{
			"status_code": http.StatusConflict,
			"data":        []Task{},
		}
	}
}

func CreateTaskProcess(title, description, dueDate string, priority int) map[string]interface{} {
	CreateTask(title, description, dueDate, priority)
	return map[string]interface{}{
		"status_code": http.StatusOK,
		"message":     "Task created",
	}
}
