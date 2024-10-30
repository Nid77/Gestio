package cli

import (
	"reflect"
	Data "gestio/data"
)

func getFields(data []Data.Task, fieldName string) []interface{} {
	var fieldValues []interface{}

	for _, task := range data {
		taskValue := reflect.ValueOf(task)
		fieldValue := taskValue.FieldByName(fieldName).Interface()
		fieldValues = append(fieldValues, fieldValue)
	}

	return fieldValues
}

func convertToStringSlice(values []interface{}) []string {
	var stringSlice []string
	for _, value := range values {
		if strValue, ok := value.(string); ok {
			stringSlice = append(stringSlice, strValue)
		}
	}
	return stringSlice
}

func getFieldNames(data interface{}) []string {
	var fieldNames []string
	dataType := reflect.TypeOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}