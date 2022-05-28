package console

import (
	"fmt"
	"github.com/gookit/event"
	"infantry/bindings"
	"log"
	"os"
)

func main() {
	fmt.Printf("Loading plan...")
	var planFile string
	var outputFile string

	err := SetupArgs(outputFile, planFile).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	event.On(bindings.StartedEvent, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Started Plan...")
		return nil
	}), event.High)

	event.On(bindings.CompletedEvent, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Complete!")
		os.Exit(0)
		return nil
	}), event.High)
}
