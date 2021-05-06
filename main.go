package main

import (
	"github.com/Jeffail/benthos/v3/lib/log"
	"github.com/Jeffail/benthos/v3/lib/metrics"
	"github.com/Jeffail/benthos/v3/lib/processor"
	"github.com/Jeffail/benthos/v3/lib/service"
	"github.com/Jeffail/benthos/v3/lib/types"
	"math/rand"
	"time"
)

func Word() string {
	words := []string{"nigga", "man", "dog", "homie", "homeboy", "playa"}
	random := rand.Intn(6)
	return words[random]
}

func Gangstaify(text string) string {
	return (text + ", " + Word())
}

type TestProcessor struct {
	prop string "OG"
}

// ProcessMessage applies the processor to a message
func (t *TestProcessor) ProcessMessage(msg types.Message) ([]types.Message, types.Response) {
	// Always create a new copy if we intend to mutate message contents.
	newMsg := msg.Copy()
	newMsg.Iter(func(i int, p types.Part) error {
		msg := p.Get()
		og := Gangstaify(string(msg))
		p.Set([]byte(og))
		return nil
	})
	return []types.Message{newMsg}, nil
}

// CloseAsync shuts down the processor and stops processing requests.
func (t *TestProcessor) CloseAsync() {
	println("Closing async")
}

// WaitForClose blocks until the processor has closed down.
func (t *TestProcessor) WaitForClose(timeout time.Duration) error {
	println("Waiting for close")
	return nil
}

func main() {
	processor.RegisterPlugin("gangstaify", func() interface{} {
		processor := &TestProcessor{
			prop: "OG",
		}
		return processor
	}, func(
		config interface{},
		manager types.Manager,
		logger log.Modular,
		metrics metrics.Type,
	) (types.Processor, error) {
		logger.Infoln("using Gangstaify version 0.1")
		return &TestProcessor{}, nil
	})
	service.Run()
}
