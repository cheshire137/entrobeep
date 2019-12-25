package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"

	"github.com/cheshire137/entrobeep/pkg/input"
	"github.com/cheshire137/entrobeep/pkg/transform"
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

	transformers := transform.GetTransformers()
	inputs := input.GetInputs()

	e := wav.NewEncoder(out, rate, bits, channels, 1)

	buf := &audio.IntBuffer{
		Format: &audio.Format{NumChannels: channels, SampleRate: rate},
	}
	for i := 0; i < 700; i++ {
		input := sampleInput(inputs)
		transformer := sampleTransformer(transformers)
		originalValue := input.Get()
		transformedValues := transformer.Transform(originalValue)
		fmt.Printf("%d => %v\n", originalValue, transformedValues)
		for _, value := range transformedValues {
			buf.Data = append(buf.Data, value)
		}
		time.Sleep(10 * time.Millisecond)
	}
	if err := e.Write(buf); err != nil {
		log.Fatal(err)
	}
	if err := e.Close(); err != nil {
		log.Fatal(err)
	}
}

func sampleInput(inputs []input.Input) input.Input {
	index := rand.Intn(len(inputs))
	return inputs[index]
}

func sampleTransformer(transformers []transform.Transformer) transform.Transformer {
	index := rand.Intn(len(transformers))
	return transformers[index]
}
