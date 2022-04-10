quad-go
=======

Небольшой фреймфорк для работы с каналами и горутинами

Usage
-----

### Generator

Принимает слайс данных, возвращает канал переданных данных 

```go
ctx, cancel := context.WithCancel(context.TODO())
defer cancel()

slice := []int{1, 2, 3}
channel := quad.Generator(ctx, slice)
for i := range channel {
	println(i)
}
```

### Batcher

```go
inputData := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
inputChannel := quad.Generator(ctx, inputData)

size := 2
batchesChannel := quad.Batcher(ctx, inputChannel, size)

for i := range batchesChannel {
	fmt.Println(i)
}
```
