package main

import "fmt"

// Task represents a single to-do item.
type Task struct {
	ID          int
	Description string
	Done        bool
}

// Complete marks the task as done.
func (t *Task) Complete() {
	t.Done = true
}

// String implements fmt.Stringer so tasks print nicely with fmt.Println etc.
func (t Task) String() string {
	status := " "
	if t.Done {
		status = "x"
	}
	return fmt.Sprintf("[%s] %d. %s", status, t.ID, t.Description)
}

// TaskList is a slice of tasks.
type TaskList []Task

// Add appends a new task with the given description and returns the updated list.
func (tl *TaskList) Add(description string) {
	*tl = append(*tl, Task{
		ID:          len(*tl) + 1,
		Description: description,
	})
}

// Complete marks the task with the given ID as done.
func (tl TaskList) Complete(id int) bool {
	for i := range tl {
		if tl[i].ID == id {
			tl[i].Complete()
			return true
		}
	}
	return false
}

// String implements fmt.Stringer for the whole list.
func (tl TaskList) String() string {
	out := ""
	for _, t := range tl {
		out += t.String() + "\n"
	}
	return out
}

func main() {
	var tasks TaskList

	tasks.Add("Write report")
	tasks.Add("Review pull request")
	tasks.Add("Buy groceries")

	tasks.Complete(2)

	fmt.Println("Task list:")
	fmt.Print(tasks)

	fmt.Println("\nSingle task:")
	fmt.Println(tasks[0])
}
