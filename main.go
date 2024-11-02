package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"

	. "gestio/cli"
	Data "gestio/data"
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
var repository Data.TaskRepository = &Data.JSONTaskRepository{FilePath: filepath}

func init() {
	rootCmd.AddCommand(addTask)
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the task")
	addTask.Flags().String("description", "", "description of the task")
	addTask.Flags().String("status", "", "status of the task (In Progress,Completed,Not Started:)")
	addTask.Flags().String("priority", "", "level of priority for a task")
	addTask.Flags().String("duedate", "", "duedate of the task")
	addTask.Flags().String("tags", "", "tags of the task")
	addTask.Flags().String("shortdescription", "", "shortdescription")

	rootCmd.AddCommand(listTask)
	listTask.Flags().Int("limit", 5, "limit the number of tasks to display")

	rootCmd.AddCommand(removeTask)
	removeTask.Flags().String("fieldname", "", "remove by fieldname")

	rootCmd.AddCommand(editTask)

	rootCmd.AddCommand(showTask)

	rootCmd.AddCommand(searchTask)
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

		err := repository.AddTask(task)
		if err != nil {
			return
		}
		fmt.Println("Added task:", task.Name)
	},
}

var listTask = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []Data.Task

		tasks, err := repository.GetAllTasks()
		if err != nil {
			return
		}

		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			limit = 5
		}

		var nameField string = "Task Name"
		var statusField string = "Status"
		var priorityField string = "Priority"

		maxNameWidth := Max(MaxLength(tasks, func(t Data.Task) string { return t.Name }), len(nameField))
		maxStatusWidth := Max(MaxLength(tasks, func(t Data.Task) string { return t.Status }), len(statusField))
		maxPriorityWidth := Max(MaxLength(tasks, func(t Data.Task) string { return t.Priority }), len(priorityField))

		// Header
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))
		fmt.Printf("| %s | %s | %s|\n", Pad(nameField, maxNameWidth), Pad(statusField, maxStatusWidth), Pad(priorityField, maxPriorityWidth+1))
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))

		// Tasks
		for index, task := range tasks {
			if index == limit {
				break
			}
			fmt.Printf("| %s | %s | %s |\n", Pad(task.Name, maxNameWidth), Pad(task.Status, maxStatusWidth), Pad(task.Priority, maxPriorityWidth))
		}

		// Footer
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", maxNameWidth+2), strings.Repeat("-", maxStatusWidth+2), strings.Repeat("-", maxPriorityWidth+2))
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

		fmt.Print("Not implemented yet")
	},
}

var searchTask = &cobra.Command{
	Use:   "search",
	Short: "search a task",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []Data.Task
		tasks, err := repository.GetAllTasks()
		if err != nil {
			return
		}

		if len(args) < 1 {
			fmt.Print("You must at least name the task to search")
			return
		}

		var foundTasks []Data.Task
		regex := regexp.MustCompile(args[0])
		for _, task := range tasks {
			if regex.MatchString(task.Name) {
				fmt.Print(task)
				foundTasks = append(foundTasks, task)
				return
			}
		}

		if len(foundTasks) == 0 {
			fmt.Print("No task found")
		} else {
			fmt.Print("Task found: ", foundTasks)
		}
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

		fmt.Print("Not implemented yet")

	},
}
