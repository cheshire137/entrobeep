package transform

type Transformer interface {
	Transform(input int) []int
}

func GetTransformers() []Transformer {
	return []Transformer{
		&NullTransformer{},
		&LengthenTransformer{},
		&MagnifyTransformer{},
		&IncrementTransformer{},
	}
}
