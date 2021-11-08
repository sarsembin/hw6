package hw6

import (
	"context"
	"errors"
	"testing"
)

func task2Good(ctx context.Context, c chan int) error {
	//fmt.Println("good")
	c <- 0
	return nil
}

func task2Error(ctx context.Context,  c chan int) error {
	err := errors.New("failed task")
	//fmt.Println("bad")
	c <- 1
	return err
}

func TestConcurrentPart2(t *testing.T) {
	totalTasks := 10
	errorTasks := 5
	tasks := make([]func(ctx context.Context, c chan int) error, totalTasks)
	for i := range tasks {
		if i < errorTasks {
			tasks[i] = task2Error
		} else {
			tasks[i] = task2Good
		}
	}
	err := Execute2(tasks, 6)
	if err != nil {
		t.Error(err)
	}

	err = Execute2(tasks, 3)
	if err == nil {
		t.Error(err)
	}

}
