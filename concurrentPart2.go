package hw6

import (
	"context"
	"errors"
)

func Execute2(tasks []func(ctx context.Context, c chan int) error, E int) error {
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, v := range tasks {
		go v(ctx, c)
	}

	cnt := 0
	total := 0
	for total < len(tasks) {
		ifErr, ok := <-c
		if !ok {
			break
		}

		cnt += ifErr
		total++

		if cnt == E {
			cancel()
			return errors.New("too many cooks")
		}
	}
	return nil
}
