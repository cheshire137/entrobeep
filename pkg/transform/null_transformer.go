package transform

type NullTransformer struct {
}

var _ Transformer = (*NullTransformer)(nil)

func (t *NullTransformer) Transform(input int) int {
	return input
}
