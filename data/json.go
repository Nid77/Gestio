package data

import (
	"encoding/json"
	"fmt"
	"os"
)

func getJsonData(filepath string) ([]Task,error) {
	var task []Task
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("File Read Error:", err)
		return nil, err
	}
	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		fmt.Println("JSON Decoding Error:", err)
		return nil, err
	}

	return task, nil
}

func saveJsonData(filepath string, task []Task) error {
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

type JSONTaskRepository struct {
    FilePath string
}

func (repo *JSONTaskRepository) AddTask(task Task) error {
    tasks, err := getJsonData(repo.FilePath)
    if err != nil {
        return err
    }
    tasks = append(tasks, task)
    return saveJsonData(repo.FilePath, tasks)
}


func (repo *JSONTaskRepository) GetAllTasks() ([]Task, error) {
    return getJsonData(repo.FilePath)
}

func (repo *JSONTaskRepository) GetTask(id int) (Task, error) {
	tasks, err := getJsonData(repo.FilePath)
	if err != nil {
		return Task{}, err
	}
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("Task with ID %d not found", id)
}

func (repo *JSONTaskRepository) UpdateTask(task Task) error {
	tasks, err := getJsonData(repo.FilePath)
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			return saveJsonData(repo.FilePath, tasks)
		}
	}
	return fmt.Errorf("Task with ID %d not found", task.ID)
}

func (repo *JSONTaskRepository) DeleteTask(id int) error {
	tasks, err := getJsonData(repo.FilePath)
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveJsonData(repo.FilePath, tasks)
		}
	}
	return fmt.Errorf("Task with ID %d not found", id)
}



