package quad

import "context"

// Batcher возвращает канал слайсов переданного размера
func Batcher[T any](ctx context.Context, in <-chan T, size int) <-chan []T {
	out := make(chan []T)

	go func() {
		defer close(out)

		batch := make([]T, 0, size)
		for item := range in {
			batch = append(batch, item)
			if len(batch) == size {
				select {
				case out <- batch:
				case <-ctx.Done():
					return
				}
				batch = make([]T, 0, size)
			}
		}
		if len(batch) > 0 {
			select {
			case out <- batch:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}
