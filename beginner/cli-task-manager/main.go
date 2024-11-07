package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

type Tasks map[int]string

func (t Tasks) GetTasks(w io.Writer) {
	ids := make([]int, 1, len(t))
	for idx := range t {
		ids = append(ids, idx)
	}

	sort.Ints(ids)
	for id := range ids {
		if id == 0 {
			continue
		}
		fmt.Fprintf(w, "%d. %s\n", id, t[id])
	}
}

func (t Tasks) Search(id int) bool {
	_, found := t[id]
	return found
}

func (t Tasks) Add(task string) {
	t[len(t)+1] = task
}

func (t Tasks) Update(id int, newTask string) error {
	isAvail := t.Search(id)
	if isAvail {
		t[id] = newTask
		return nil
	}
	return errors.New("couldnt update,task doesnt exist")
}

func (t Tasks) Delete(id int) error {
	isAvail := t.Search(id)
	if isAvail {
		delete(t, id)
		return nil
	}
	return errors.New("couldnt delete,task doesnt exist")
}

func main() {
	var tasks Tasks
	tasks = make(map[int]string)
	tasks.Add("watch")
	tasks.Add("clean")
	tasks.Add("draw")

	tasks.GetTasks(os.Stdout)
}
