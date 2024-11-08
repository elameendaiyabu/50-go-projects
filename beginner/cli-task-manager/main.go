package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Tasks map[int]string

func (t Tasks) GetTasks(w io.Writer) {
	if len(t) < 1 {
		return
	}
	ids := make([]int, 1, len(t))
	for idx := range t {
		ids = append(ids, idx)
	}

	sort.Ints(ids)
	for id, key := range ids {
		if id == 0 {
			continue
		}
		fmt.Fprintf(w, "%d. %s\n", key, t[key])
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

func (t *Tasks) Delete(id int) error {
	isAvail := t.Search(id)
	if isAvail {
		delete(*t, id)
		return nil
	}
	return errors.New("couldnt delete,task doesnt exist")
}

func main() {
	var tasks Tasks
	tasks = make(map[int]string)

	for {
		fmt.Print("\033[H\033[2J")
		tasks.GetTasks(os.Stdout)

		fmt.Println("enter a command")
		var cmd string
		fmt.Scanln(&cmd)

		if cmd == "q" {
			os.Exit(1)
		}

		if cmd == "add" || cmd == "a" {
			fmt.Println("Enter a new task")
			reader := bufio.NewReader(os.Stdin)
			task, err := reader.ReadString('\n')
			handleErr(err)
			tasks.Add(task)
		}

		if cmd == "delete" || cmd == "d" {
			fmt.Println("enter task id to delete")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err := scanner.Err()
			handleErr(err)
			taskId, _ := strconv.Atoi(scanner.Text())
			tasks.Delete(int(taskId))
		}

		if cmd == "update" || cmd == "u" {
			fmt.Println("enter task id to update")
			id := bufio.NewScanner(os.Stdin)
			id.Scan()
			err := id.Err()
			handleErr(err)
			taskId, _ := strconv.Atoi(id.Text())
			fmt.Println("enter new task to update")
			task := bufio.NewScanner(os.Stdin)
			task.Scan()
			err = task.Err()
			handleErr(err)
			tasks.Update(int(taskId), task.Text())
		}

	}
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
