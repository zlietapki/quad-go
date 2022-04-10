package quad

import (
	"context"
	"sync"
)

type Worker[inT any, outT any] interface {
	func(ctx context.Context, in <-chan inT, out chan<- outT, errChan chan<- error)
}

func MakeWorkerPool[inT any, outT any, workerT Worker[inT, outT]](ctx context.Context, worker workerT, size int, in <-chan inT) (<-chan outT, <-chan error) {
	out := make(chan outT)
	errChan := make(chan error)

	wg := new(sync.WaitGroup)
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, in, out, errChan)
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		close(errChan)
	}()

	return out, errChan
}
