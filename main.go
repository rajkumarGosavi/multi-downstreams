package main

import (
	"abstraction/models"
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

	analyticRespChan := make(chan models.AnalyseChanData)
	feedbackRespChan := make(chan models.FeedbackChanData)
	AnalysisDone := make(chan struct{})
	FeedbackDone := make(chan struct{})

	for _, e := range events {
		switch e.Action {
		case "analyse":
			go func() {
				for {
					select {
					case <-AnalysisDone:
						return
					case data := <-analyticRespChan:
						m.SaveAnalyserData(data)
					}
				}
			}()
			m.AnalyseLoop(e, analyticRespChan, AnalysisDone)

		case "feedback":
			go func() {
				for {
					select {
					case <-FeedbackDone:
						return
					case data := <-feedbackRespChan:
						m.SaveFeedbackerData(data)
					}
				}
			}()

			m.FeedbackLoop(e, feedbackRespChan, FeedbackDone)
		}
	}

	fmt.Println("!!! Done")
}
