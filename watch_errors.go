package quad

import "context"

func WatchErrors(errs <-chan error, errorVar *error, cancel context.CancelFunc) {
	go func() {
		for err := range errs {
			if errorVar != nil {
				*errorVar = err
			}
			if cancel != nil {
				cancel()
			}
		}
	}()
}
