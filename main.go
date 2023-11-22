package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/dmsyudha/task-scheduler/scheduler"
)

func main() {
	s := scheduler.NewScheduler()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Add task")
		fmt.Println("2. Remove task")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')

		switch choice {
		case "1\n":
			fmt.Print("Enter task ID: ")
			id, _ := reader.ReadString('\n')

			fmt.Print("Enter task name: ")
			name, _ := reader.ReadString('\n')

			fmt.Print("Enter task execution time (in seconds): ")
			var execTime int
			fmt.Scanf("%d", &execTime)

			t := scheduler.NewTask(id, name, time.Now().Add(time.Duration(execTime)*time.Second), func() {
				fmt.Println("Executing task", id)
			})

			s.AddTask(t)
		case "2\n":
			fmt.Print("Enter task ID to remove: ")
			id, _ := reader.ReadString('\n')

			s.RemoveTask(id)
		case "3\n":
			return
		}
	}
}
