package main

import "fmt"

type Tasks map[int]string

func (t Tasks) GetTasks() {
	for idx, task := range t {
		fmt.Printf("%d. %s\n", idx, task)
	}
}

func (t Tasks) Search(id int) bool {
	_, found := t[id]
	return found
}

func (t Tasks) Add(task string) {
	t[len(t)+1] = task
}

func (t Tasks) Update(id int, newTask string) {
	t[id] = newTask
}

func (t Tasks) Delete(id int) {
	delete(t, id)
}

func main() {
	var tasks Tasks
	tasks = make(map[int]string)
	tasks.Add("watch")
	tasks.Add("clean")
	tasks.Add("draw")

	tasks.GetTasks()
}
