package quad

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBatch(t *testing.T) {
	ctx := context.Background()

	t.Run("[]int", func(t *testing.T) {
		genChannel := Generator(ctx, []int{1, 2, 3})
		batchChannel := Batcher(ctx, genChannel, 2)
		expected := [][]int{
			{1, 2},
			{3},
		}

		res := make([][]int, 0)
		for val := range batchChannel {
			res = append(res, val)
		}
		require.Equal(t, expected, res)
	})

	t.Run("[]int64", func(t *testing.T) {
		genChannel := Generator(ctx, []int64{1, 2, 3})
		batchChannel := Batcher(ctx, genChannel, 2)
		expected := [][]int64{
			{1, 2},
			{3},
		}

		res := make([][]int64, 0)
		for val := range batchChannel {
			res = append(res, val)
		}
		require.Equal(t, expected, res)
	})

	t.Run("[]uint64", func(t *testing.T) {
		genChannel := Generator(ctx, []uint64{1, 2, 3})
		batchChannel := Batcher(ctx, genChannel, 2)
		expected := [][]uint64{
			{1, 2},
			{3},
		}

		res := make([][]uint64, 0)
		for val := range batchChannel {
			res = append(res, val)
		}
		require.Equal(t, expected, res)
	})
}
