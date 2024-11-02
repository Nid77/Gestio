package data

type Task struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortdescription"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	Priority         string `json:"priority"`
	DueDate          string `json:"duedate"`
	Tags             string `json:"tags"`
	CreationDate     string `json:"creation_date"`
}

type TaskRepository interface {
    AddTask(task Task) error
    GetTask(id int) (Task, error)
    GetAllTasks() ([]Task, error)
    UpdateTask(task Task) error
    DeleteTask(id int) error
}