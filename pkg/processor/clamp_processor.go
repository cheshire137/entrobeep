package processor

import (
	"modernc.org/mathutil"
)

type ClampProcessor struct {
}

var _ Processor = (*ClampProcessor)(nil)

func (p *ClampProcessor) Process(value int) int {
	return mathutil.Clamp(value, 50000, 100000000)
}
