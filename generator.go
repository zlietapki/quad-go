package quad

import "context"

// Generator возвращает канал переданных параметров
func Generator[T any](ctx context.Context, slice []T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for _, item := range slice {
			select {
			case out <- item:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}
