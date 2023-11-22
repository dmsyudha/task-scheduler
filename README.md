# Task Scheduler in GoLang

This project is a simple task scheduler implemented in GoLang. The scheduler allows users to add, remove, and execute tasks at specified times.

## Project Structure

The project is structured into two main files:

1. `main.go`: This file contains the main function and handles user interaction.
2. `scheduler/task.go` and `scheduler/scheduler.go`: These files define the `Task` and `Scheduler` structs and their associated methods.

## Usage

To run the task scheduler, execute the `main.go` file with the `go run` command:

```bash
go run main.go
```

The program will present a simple command-line interface that allows you to add and remove tasks. When adding a task, you will be prompted to enter the task ID, name, and execution time. The task function simply prints a message when the task is executed. When removing a task, you will be prompted to enter the task ID of the task to remove.

## Function Plugins

Users can provide their own functions for tasks in separate Go files in the `functions` directory. Each function should be in a file with the `.go` extension and have the signature `func()`. The function should be named `Function` and the package should be `main`.

To build a function plugin, use the `go build -buildmode=plugin` command. For example, to build a function in a file named `hello.go`, use the following command:

```bash
go build -buildmode=plugin -o functions/hello.so functions/hello.go
```

Then, when adding a task, enter `hello.so` as the function file name.

## Testing

The `scheduler_test.go` file contains tests for the `Scheduler` struct. You can run the tests by executing the `scheduler_test.go` file with the `go test` command:

```bash
go test -timeout 30s -run ^TestScheduler$ github.com/dmsyudha/task-scheduler/scheduler
```

The tests check that tasks are correctly added and removed from the scheduler.