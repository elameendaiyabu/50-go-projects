package main

import "fmt"

type Task map[int]string

func PrintAllTasks(tasks map[int]string) {
	for id, task := range tasks {
		fmt.Printf("%d. %s", task[id], task)
	}
}

func (t Task) Add(task string) {
	t[len(t)+1] = task
}

func (t Task) GetTasks() {
	for idx, task := range t {
		fmt.Printf("%d. %s\n", idx, task)
	}
}

func main() {
	var tasks Task
	tasks = make(map[int]string)
	tasks.Add("watch")
	tasks.Add("clean")
	tasks.Add("draw")

	tasks.GetTasks()
}
