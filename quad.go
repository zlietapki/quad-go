package quad

import (
	"context"
	"sync"
)

type Object struct {
	Value interface{}
	Error error
}

type Worker[inT any, outT any] interface {
	func(ctx context.Context, in <-chan inT, out chan<- outT)
}

func MakeWorkerPool[inT any, outT any, workerT Worker[inT, outT]](ctx context.Context, worker workerT, size int, in <-chan inT) <-chan outT {
	out := make(chan outT)

	wg := new(sync.WaitGroup)
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, in, out)
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
