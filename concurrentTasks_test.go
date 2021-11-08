package hw6

import (
	"errors"
	"testing"
)

func taskGood(c chan error) error {
	c <- nil
	//fmt.Println("good")
	return nil
}

func taskError(c chan error) error {
	err := errors.New("failed task")
	c <- err
	//fmt.Println("bad")
	return err
}

func TestConcurrentTasks(t *testing.T) {
	totalTasks := 10
	errorsCount :=  6
	tasks := make([]func(c chan error) error, totalTasks)
	for i := range tasks{
		if i < errorsCount {
			tasks[i] = taskError
		} else {
			tasks[i] = taskGood
		}
	}
	err := Execute(tasks, 5)
	if err == nil {
		t.Error(err)
	}

	err = Execute(tasks, 7)
	if err != nil {
		t.Error(err)
	}
}
