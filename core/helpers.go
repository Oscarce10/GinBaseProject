package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
)

// ToString Change arg to string
func ToString(arg interface{}, timeFormat ...string) string {
	if len(timeFormat) > 1 {
		log.SetFlags(log.Llongfile | log.LstdFlags)
		log.Println(errors.New(fmt.Sprintf("timeFormat's length should be one")))
	}
	var tmp = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	switch v := tmp.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		if len(timeFormat) == 1 {
			return v.Format(timeFormat[0])
		}
		return v.Format("2006-01-02 03:04:05")
	case bool:
		if arg.(bool) {
			return "true"
		} else {
			return "false"
		}
	case interface{}:
		mJson, err := json.Marshal(arg)
		if err != nil {
			return "ERROR"
		}
		return string(mJson)
	default:
		return ""
	}
}
