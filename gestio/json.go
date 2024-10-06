package main

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

func Write(filepath string, task *[]Task) {
	jsonData, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		fmt.Println("Erreur d'encodage JSON:", err)
		return
	}

	err = os.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture dans le fichier :", err)
		return
	}
}

func Open(filepath string, task *[]Task) {
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		fmt.Println("Erreur de décodage JSON:", err)
		return
	}
}
