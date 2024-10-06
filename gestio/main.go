package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gestio",
	Short: "A simple task manager CLI app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome !")
	},
}
var name string

func init() {
	rootCmd.AddCommand(addTask)
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the task")
	addTask.Flags().String("description", "", "description of the task")
	addTask.Flags().String("status", "", "status of the task (In Progress,Completed,Not Started:)")
	addTask.Flags().String("priority", "", "level of priority for a task")
	addTask.Flags().String("duedate", "", "duedate of the task")
	addTask.Flags().String("tags", "", "tags of the task")
	addTask.Flags().String("shortdescription", "", "shortdescription")
	rootCmd.AddCommand(removeTask)
	removeTask.Flags().String("fieldname", "", "remove by fieldname")
	rootCmd.AddCommand(editTask)
	rootCmd.AddCommand(listTask)
	rootCmd.AddCommand(showTask)
}

var addTask = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {

		if (name == "") && (args[0] == "") {
			fmt.Print("You must at least name the task to add")
			return
		}

		if name == "" {
			name = args[0]
		}

		var tasks []Task
		Open("data.json", &tasks)
		task := Task{
			Name:             name,
			ShortDescription: cmd.Flags().Lookup("shortdescription").Value.String(),
			Description:      cmd.Flags().Lookup("description").Value.String(),
			CreationDate:     time.Now().Format("2006-01-02"),
			Status:           cmd.Flags().Lookup("status").Value.String(),
			Priority:         cmd.Flags().Lookup("priority").Value.String(),
			DueDate:          cmd.Flags().Lookup("duedate").Value.String(),
			Tags:             cmd.Flags().Lookup("tags").Value.String(),
		}

		tasks = append(tasks, task)

		Write("data.json", &tasks)

		fmt.Println("Added task:", task.Name)
	},
}

var removeTask = &cobra.Command{
	Use:   "rm",
	Short: "remove a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 && name == "" {
			fmt.Print("You must at least name the task to remove")
			return
		}

		if name == "" && len(args) > 1 {
			name = args[0]
		}

		var tasks []Task
		Open("data.json", &tasks)
		for _, task := range tasks {
			if task.Name == name {
				fmt.Print(task)

				return
			}
		}
		fmt.Print("this task don't exist")

	},
}

var editTask = &cobra.Command{
	Use:   "edit",
	Short: "modify a task",
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			fmt.Print("You must at least name the task to edit")
			return
		}
		var tasks []Task
		Open("data.json", &tasks)

	},
}

var showTask = &cobra.Command{
	Use:   "show",
	Short: "show a specefic detail of a task",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []Task
		Open("data.json", &tasks)
		for _, task := range tasks {
			if task.Name == args[0] {
				fmt.Print(task)
				return
			}
		}
		fmt.Print("this task don't exist")

	},
}

var listTask = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []Task
		//var str [8][]string

		//fieldName := []string{"Name", "Status", "Priority"} //getFieldNames(Task{})
		//max := make([]int, len(fieldName))
		//o := 0
		Open("data.json", &tasks)
		/*
			for _, name := range fieldName {
				fmt.Print(name + " : ")
				fmt.Print(convertToStringSlice(extractField(tasks, name)))
				fmt.Print("\n")
			}
		*/
		maxNameWidth := maxLength(tasks, func(t Task) string { return t.Name })
		maxStatusWidth := maxLength(tasks, func(t Task) string { return t.Status })
		maxPriorityWidth := maxLength(tasks, func(t Task) string { return t.Priority })

		// En-tête du tableau
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))
		fmt.Printf("| %s | %s | %s|\n", pad("Nom de la tâche", maxNameWidth), pad("Statut", maxStatusWidth), pad("Priorité", maxPriorityWidth))
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))

		// Affichage des tâches dans un tableau
		for _, task := range tasks {
			fmt.Printf("| %s | %s | %s |\n", pad(task.Name, maxNameWidth), pad(task.Status, maxStatusWidth), pad(task.Priority, maxPriorityWidth))
		}

		// Ligne inférieure du tableau
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))
	},
}

func pad(str string, length int) string {
	return fmt.Sprintf("%-*s", length, str)
}

func maxLength(tasks []Task, selector func(Task) string) int {
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

func extractField(data []Task, fieldName string) []interface{} {
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
