package input

type Input interface {
	Get() int
}

func GetInputs() []Input {
	return []Input{
		&MouseInput{},
	}
}