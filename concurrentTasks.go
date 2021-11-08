package hw6

import (
	"errors"
)


func Execute(tasks []func(c chan error) error, E int) error {
	c := make(chan error)
	done := make(chan int)
	for _, v := range tasks {
		go v(c)
	}
	go receiver(c, done, E, len(tasks))

	if errorCount := <- done; errorCount < E {
		return nil
	}
	return errors.New("too many cooks")
}

func receiver(c chan error, done chan int, E, lenAll int) {
	errCnt, total := 0, 0
	for{
		err, ok := <- c
		if !ok {
			break
		}

		if err != nil {
			errCnt++
		}
		total++

		if errCnt == E  || total == lenAll {
			done <- errCnt
			close(done)
			return
		}
	}
}



