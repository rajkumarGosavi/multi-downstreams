package main

import (
	"abstraction/monitor"
	"fmt"
)

func main() {
	events := []struct {
		Action string
	}{
		{"analyse"},
		{"analyse"},
		{"feedback"},
		{"feedback"},
	}

	m := monitor.New()

	for _, e := range events {
		switch e.Action {
		case "analyse":
			m.Analyse(e)
		case "feedback":
			m.Feedback(e)
		}
	}

	fmt.Println("!!! Done")
}
