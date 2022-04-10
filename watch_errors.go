package quad

import (
	"context"
	"sync"
)

func WatchErrors(errorsChannel <-chan error, wg *sync.WaitGroup, errorVar *error, cancel context.CancelFunc) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range errorsChannel {
			if errorVar != nil {
				*errorVar = err
			}
			if cancel != nil {
				cancel()
			}
		}
	}()
}
