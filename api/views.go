package api

import (
	"GinBaseProject/core"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func rootView(c *gin.Context) {
	c.JSONP(
		http.StatusOK,
		gin.H{
			"message": "Base api for Go built with Gin",
		},
	)
}

func GetTasksView(c *gin.Context) {
	res := core.GetTasksProcess()
	status := res["status_code"].(int)
	delete(res, "status_code")
	c.JSONP(
		status,
		res,
	)
	return
}

func CreateTaskView(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.JSONP(
			http.StatusUnsupportedMediaType,
			gin.H{
				"message": "Content-Type must be application/json",
			},
		)
		return
	}
	request := core.CreateTaskInputValidator{}
	if err := c.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]core.ErrorMsg, len(ve))
			for i, err := range ve {
				out[i] = core.ErrorMsg{
					Field:   err.Field(),
					Message: core.GetErrorMsg(err),
				}
			}
			c.JSONP(
				http.StatusBadRequest,
				gin.H{
					"data": out,
				})
			return
		} else {
			c.JSONP(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
				},
			)
			return
		}
	}
	title := core.ToString(request.Title)
	description := request.Description
	dueDate := request.DueDate.String()
	priority := request.Priority

	res := core.CreateTaskProcess(title, description, dueDate, priority)
	status := res["status_code"].(int)
	delete(res, "status_code")
	c.JSONP(
		status,
		res,
	)
	return
}
