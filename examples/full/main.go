package main

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/zlietapki/quad-go"
)

type MyData struct {
	Name string
	Val  int
}

func worker1(ctx context.Context, in <-chan []MyData, out chan<- string, err chan error) {
	for batch := range in {
		for _, myData := range batch {
			if myData.Name == "incorrect name" {
				select {
				case err <- errors.New("some worker1 error"):
				case <-ctx.Done():
					return
				}
				continue
			}

			select {
			case out <- myData.Name:
			case <-ctx.Done():
				return
			}
		}
	}
}

func worker2(ctx context.Context, in <-chan string, out chan<- string, err chan error) {
	for name := range in {
		if name == "Andrey" {
			select {
			case err <- errors.New("worker 2 error"):
			case <-ctx.Done():
				return
			}
		}

		select {
		case out <- name + " patched":
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := []MyData{
		{
			Name: "Andrey",
			Val:  111,
		},
		{
			Name: "Alexey",
			Val:  222,
		},
		{
			Name: "incorrect name",
			Val:  333,
		},
	}

	// data > generator > batcher > worker1 > worker2
	// profit

	inputChannel := quad.Generator(ctx, data)
	batchesChannel := quad.Batcher(ctx, inputChannel, 2)

	var myerror error

	res1, err1 := quad.MakeWorkerPool(ctx, worker1, 2, batchesChannel)
	wgErr := new(sync.WaitGroup)
	quad.WatchErrors(err1, wgErr, &myerror, nil)

	res2, err2 := quad.MakeWorkerPool(ctx, worker2, 2, res1)
	quad.WatchErrors(err2, wgErr, &myerror, nil)

	for i := range res2 {
		fmt.Println(i)
	}
	wgErr.Wait()

	if myerror != nil {
		fmt.Println(myerror)
	}
}
