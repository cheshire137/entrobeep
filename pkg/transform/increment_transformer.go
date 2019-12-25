package transform

import "math/rand"

type IncrementTransformer struct {
}

var _ Transformer = (*IncrementTransformer)(nil)

func (t *IncrementTransformer) Transform(input int) []int {
	times := rand.Intn(30)
	result := make([]int, times)
	for i := 0; i < times; i++ {
		result[i] = input + i
	}
	return result
}
