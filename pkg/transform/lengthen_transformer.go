package transform

import "math/rand"

type LengthenTransformer struct {
}

var _ Transformer = (*LengthenTransformer)(nil)

func (t *LengthenTransformer) Transform(input int) []int {
	times := rand.Intn(30)
	result := make([]int, times)
	for i := 0; i < times; i++ {
		result[i] = input
	}
	return result
}
