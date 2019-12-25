package input

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

type MouseInput struct {
}

var _ Input = (*MouseInput)(nil)

func (i *MouseInput) Get() int {
	x, y := robotgo.GetMousePos()
	randomNum := rand.Intn(4)
	if randomNum < 1 {
		return x
	}
	if randomNum < 2 {
		return y
	}
	if randomNum < 3 {
		return x + y
	}
	return x * y // if randomNum == 3
}
