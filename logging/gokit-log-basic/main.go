package main

import (
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	type Task struct {
		ID int
	}

	RunTask := func(task Task, logger log.Logger) {
		logger.Log("taskID", task.ID, "event", "starting task")

		logger.Log("taskID", task.ID, "event", "task complete")
	}

	RunTask(Task{ID: 1}, logger)
}
