package transform

import "math/rand"

type MagnifyTransformer struct {
}

var _ Transformer = (*MagnifyTransformer)(nil)

func (t *MagnifyTransformer) Transform(input int) []int {
	value := input * rand.Intn(10)
	return []int{value}
}
