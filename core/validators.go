package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"time"
)

type onlyDateFormat time.Time

var _ json.Unmarshaler = &onlyDateFormat{}

func (t *onlyDateFormat) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	match, _ := regexp.MatchString(`^[1-2][0-9]{3}(?:-[0-9]{2}){2}$`, s)
	if !match {
		return errors.New(fmt.Sprintf("Date provided (%s) is not a valid date", s))
	}
	tt, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*t = onlyDateFormat(tt)
	return nil
}

func (t *onlyDateFormat) String() string {
	return time.Time(*t).Format("2006-01-02")
}

type CreateTaskInputValidator struct {
	Title       string          `form:"title" binding:"required"`
	Description string          `form:"description" binding:"required"`
	Priority    int             `form:"priority" binding:"required,gte=1,lte=5"`
	DueDate     *onlyDateFormat `json:"due_date" binding:"required" time_format:"2006-01-02"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(param validator.FieldError) string {
	switch param.Tag() {
	case "required":
		return "The field is required"
	case "gte":
		return "The field must be greater than or equal to " + param.Param()
	case "lte":
		return "The field must be less than or equal to " + param.Param()
	default:
		return "The field is invalid"
	}
}
