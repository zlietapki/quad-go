package main

import (
	"context"
	"fmt"

	"github.com/zlietapki/quad-go"
)

type MyData struct {
	Name string
	Val  int
}

func worker(ctx context.Context, in <-chan []MyData, out chan<- string) {
	for batch := range in {
		for _, myData := range batch {
			select {
			case out <- myData.Name:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx := context.Background()

	data := []MyData{
		{
			Name: "name 1",
			Val:  111,
		},
		{
			Name: "name 2",
			Val:  222,
		},
		{
			Name: "name 3",
			Val:  333,
		},
	}

	inputChannel := quad.Generator(ctx, data)
	batchesChannel := quad.Batcher(ctx, inputChannel, 2)
	resultsChannel := quad.MakeWorkerPool(ctx, worker, 2, batchesChannel)

	for i := range resultsChannel {
		fmt.Println(i)
	}
}
