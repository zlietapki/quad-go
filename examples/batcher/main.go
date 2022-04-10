package main

import (
	"context"
	"fmt"

	"github.com/zlietapki/quad-go"
)

func main() {
	ctx := context.Background()

	inputData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputChannel := quad.Generator(ctx, inputData)

	size := 4
	batchesChannel := quad.Batcher(ctx, inputChannel, size)

	for i := range batchesChannel {
		fmt.Println(i)
	}
}
