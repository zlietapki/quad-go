package quad

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	ctx := context.Background()

	t.Run("[]int", func(t *testing.T) {
		items := []int{1, 2, 3}

		ch := Generator(ctx, items)
		res := make([]int, 0, len(items))
		for i := range ch {
			res = append(res, i)
		}
		require.Equal(t, items, res)
	})

	t.Run("[]uint64", func(t *testing.T) {
		items := []uint64{1, 2, 3}

		ch := Generator(ctx, items)
		res := make([]uint64, 0, len(items))
		for i := range ch {
			res = append(res, i)
		}
		require.Equal(t, items, res)
	})

	t.Run("[][]int", func(t *testing.T) {
		items := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		ch := Generator(ctx, items)
		res := make([][]int, 0, len(items))
		for i := range ch {
			res = append(res, i)
		}
		require.Equal(t, items, res)
	})
}
