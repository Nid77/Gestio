package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	Data "gestio/data"
	//Cli "gestio/cli"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gestio",
	Short: "A simple task manager CLI app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome !")
	},
}

var name string
var filepath string = "data/data.json"

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

		task := Data.Task{
			Name:             name,
			ShortDescription: cmd.Flags().Lookup("shortdescription").Value.String(),
			Description:      cmd.Flags().Lookup("description").Value.String(),
			CreationDate:     time.Now().Format("2006-01-02"),
			Status:           cmd.Flags().Lookup("status").Value.String(),
			Priority:         cmd.Flags().Lookup("priority").Value.String(),
			DueDate:          cmd.Flags().Lookup("duedate").Value.String(),
			Tags:             cmd.Flags().Lookup("tags").Value.String(),
		}

		var tasks []Data.Task
		if err := Data.GetJsonData(filepath, &tasks); err != nil {
			return
		}

		tasks = append(tasks, task)

		if err := Data.SaveJsonData(filepath, &tasks); err != nil {
			return
		}

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

		var tasks []Data.Task

		if err := Data.GetJsonData(filepath, &tasks); err != nil {
			return
		}

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

		fmt.Print("Not implemented yet")
	},
}

var showTask = &cobra.Command{
	Use:   "show",
	Short: "show a specefic detail of a task",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []Data.Task

		if err := Data.GetJsonData(filepath, &tasks); err != nil {
			return
		}

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
		var tasks []Data.Task

		if err := Data.GetJsonData(filepath, &tasks); err != nil {
			return
		}

		maxNameWidth := maxLength(tasks, func(t Data.Task) string { return t.Name })
		maxStatusWidth := maxLength(tasks, func(t Data.Task) string { return t.Status })
		maxPriorityWidth := maxLength(tasks, func(t Data.Task) string { return t.Priority })

		// Header
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))
		fmt.Printf("| %s | %s | %s|\n", pad("Nom de la tâche", maxNameWidth), pad("Statut", maxStatusWidth), pad("Priorité", maxPriorityWidth))
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))

		// Tasks
		for _, task := range tasks {
			fmt.Printf("| %s | %s | %s |\n", pad(task.Name, maxNameWidth), pad(task.Status, maxStatusWidth), pad(task.Priority, maxPriorityWidth))
		}

		// Footer
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))
	},
}

func pad(str string, length int) string {
	return fmt.Sprintf("%-*s", length, str)
}

func maxLength(tasks []Data.Task, selector func(Data.Task) string) int {
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
