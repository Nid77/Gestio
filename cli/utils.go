package cli

import (
	"fmt"
	Data "gestio/data"
	"reflect"
)

func Pad(str string, length int) string {
	return fmt.Sprintf("%-*s", length, str)
}

func MaxLength(tasks []Data.Task, selector func(Data.Task) string) int {
	max := 0
	for _, task := range tasks {
		length := len(selector(task))
		if length > max {
			max = length
		}
	}
	return max
}

func MaxString(str []string) int {
	new := ""
	for _, s := range str {
		if len(s) > len(new) {
			new = s
		}
	}
	return len(new)
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetFields(data []Data.Task, fieldName string) []interface{} {
	var fieldValues []interface{}

	for _, task := range data {
		taskValue := reflect.ValueOf(task)
		fieldValue := taskValue.FieldByName(fieldName).Interface()
		fieldValues = append(fieldValues, fieldValue)
	}

	return fieldValues
}

func ConvertToStringSlice(values []interface{}) []string {
	var stringSlice []string
	for _, value := range values {
		if strValue, ok := value.(string); ok {
			stringSlice = append(stringSlice, strValue)
		}
	}
	return stringSlice
}

func GetFieldNames(data interface{}) []string {
	var fieldNames []string
	dataType := reflect.TypeOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}
