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
	func(ctx context.Context, in <-chan inT, out chan<- outT, err chan error)
}

func MakeWorkerPool[inT any, outT any, workerT Worker[inT, outT]](ctx context.Context, worker workerT, size int, in <-chan inT) (<-chan outT, chan error) {
	out := make(chan outT)
	err := make(chan error)

	wg := new(sync.WaitGroup)
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, in, out, err)
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		close(err)
	}()

	return out, err
}
