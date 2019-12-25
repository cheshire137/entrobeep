package input

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

type DisplayInput struct {
}

var _ Input = (*DisplayInput)(nil)

func (i *DisplayInput) Get() int {
	randomNum := rand.Intn(3)
	if randomNum < 1 {
		return robotgo.Scale()
	}
	if randomNum < 2 {
		return robotgo.ScaleX()
	}
	return robotgo.ScaleY()
}
