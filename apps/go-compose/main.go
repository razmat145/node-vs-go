package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/razmat145/learning/go-compose/bubble"
	"github.com/razmat145/learning/go-compose/builder"
)

func main() {
	inChan := make(chan builder.BuildProgress, 1)
	outChan := make(chan builder.RunEnvironment, 1)

	go func() {
		runEnv := <-outChan
		inChan <- builder.BuildProgress{Progress: "building"}
		runEnv.Build()
		inChan <- builder.BuildProgress{Progress: "done"}
	}()

	p := tea.NewProgram(bubble.InitialiseModel(inChan, outChan))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
