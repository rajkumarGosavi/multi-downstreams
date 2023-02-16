package monitor

import (
	"abstraction/models"
	"abstraction/providera"
	"abstraction/providerb"
	"fmt"
)

type Monitor struct {
	analysers   []models.Analyser
	feedbackers []models.Feedbacker
	app         interface{}
}

func (m *Monitor) AnalyseLoop(event interface{}, ac chan models.AnalyseChanData, done chan struct{}) {
	for _, a := range m.analysers {
		req := models.StdAnalyseReqBody{}
		resp := a.Analyse(req)
		ac <- models.AnalyseChanData{
			Req:  req,
			Resp: resp,
		}
	}
	done <- struct{}{}
}

func (m *Monitor) FeedbackLoop(event interface{}, fc chan models.FeedbackChanData, done chan struct{}) {
	for _, a := range m.feedbackers {
		req := models.StdFeedbackReqBody{}
		resp := a.Feedback(req)
		fc <- models.FeedbackChanData{
			Req:  req,
			Resp: resp,
		}
	}
	done <- struct{}{}
}

func (m *Monitor) SaveAnalyserData(data models.AnalyseChanData) error {
	fmt.Println("Saved analyser data", data)
	return nil
}

func (m *Monitor) SaveFeedbackerData(data models.FeedbackChanData) error {
	fmt.Println("Saved feedbacker data", data)
	return nil
}

func New() *Monitor {
	b := providerb.New()
	a := providera.New()

	return &Monitor{
		analysers:   []models.Analyser{b, a},
		feedbackers: []models.Feedbacker{b, a},
		app:         nil,
	}
}
