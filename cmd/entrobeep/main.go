package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/go-vgo/robotgo"
)

func main() {
	bits := 16
	rate := 8000 // 8 kHz
	channels := 1

	out, err := os.Create("test.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	e := wav.NewEncoder(out, rate, bits, channels, 1)

	buf := &audio.IntBuffer{
		Format: &audio.Format{NumChannels: channels, SampleRate: rate},
	}
	for i := 0; i < 700; i++ {
		x, y := robotgo.GetMousePos()
		fmt.Printf("%d, %d\n", x, y)
		buf.Data = append(buf.Data, x)
		buf.Data = append(buf.Data, y)
		time.Sleep(40 * time.Millisecond)
	}
	if err := e.Write(buf); err != nil {
		log.Fatal(err)
	}
	if err := e.Close(); err != nil {
		log.Fatal(err)
	}
}
