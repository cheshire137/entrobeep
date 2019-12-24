package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	x, y := robotgo.GetMousePos()
	fmt.Printf("x: %d, y: %d\n", x, y)
}
