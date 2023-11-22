package main

import (
	"bufio"
	"fmt"
	"os"
	"plugin"
	"strconv"
	"strings"
	"time"

	"github.com/dmsyudha/task-scheduler/scheduler"
)

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getFunction(funcFile string) (func(), error) {
	p, err := plugin.Open("./functions/" + funcFile + ".so")
	if err != nil {
		return nil, err
	}

	f, err := p.Lookup("Function")
	if err != nil {
		return nil, err
	}

	function, ok := f.(func())
	if !ok {
		return nil, fmt.Errorf("function has wrong signature")
	}

	return function, nil
}

func main() {
	s := scheduler.NewScheduler()

	for {
		fmt.Println("1. Add task")
		fmt.Println("2. Remove task")
		fmt.Println("3. Exit")

		switch getUserInput("Enter choice: ") {
		case "1":
			id := getUserInput("Enter task ID: ")
			name := getUserInput("Enter task name: ")
			execTime, _ := strconv.Atoi(getUserInput("Enter task execution time (in seconds): "))
			funcFile := getUserInput("Enter function file name: ")

			function, err := getFunction(funcFile)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			t := scheduler.NewTask(id, name, time.Now().Add(time.Duration(execTime)*time.Second), function)

			err = s.AddTask(t)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "2":
			id := getUserInput("Enter task ID to remove: ")
			s.RemoveTask(id)
		case "3":
			return
		default:
			fmt.Println("Error: Invalid choice")
		}
	}
}
