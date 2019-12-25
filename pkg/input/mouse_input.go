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
	randomNum := rand.Intn(2)
	if randomNum < 1 {
		return x * 10000
	}
	return y * 10000
}
