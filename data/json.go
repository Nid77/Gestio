package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Name             string `json:"name"`
	ShortDescription string `json:"shortdescription"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	Priority         string `json:"priority"`
	DueDate          string `json:"duedate"`
	Tags             string `json:"tags"`
	CreationDate     string `json:"creation_date"`
}

func GetJsonData(filepath string, task *[]Task) error {
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("File Read Error:", err)
		return err
	}
	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		fmt.Println("JSON Decoding Error:", err)
		return err
	}

	return nil
}


func SaveJsonData(filepath string, task *[]Task) error {
	jsonData, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		fmt.Println("JSON Encoding Error:", err)
		return err
	}

	err = os.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		fmt.Println("File Write Error:", err)
		return err
	}

	return nil
}

