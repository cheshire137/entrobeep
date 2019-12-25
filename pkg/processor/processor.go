package processor

type Processor interface {
	Process(value int) int
}

func GetProcessors() []Processor {
	return []Processor{
		&ClampProcessor{},
	}
}
