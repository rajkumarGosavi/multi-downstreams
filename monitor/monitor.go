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

func (m *Monitor) Analyse(event interface{}) []models.StdAnalyseRespBody {
	var res []models.StdAnalyseRespBody
	for _, a := range m.analysers {
		req := models.StdAnalyseReqBody{}
		resp := a.Analyse(req)
		res = append(res, resp)

		m.SaveAnalyserData(req, resp)
	}

	return res
}

func (m *Monitor) Feedback(event interface{}) []models.StdFeedbackRespBody {
	var res []models.StdFeedbackRespBody
	for _, a := range m.feedbackers {
		req := models.StdFeedbackReqBody{}
		resp := a.Feedback(req)
		res = append(res, resp)
		m.SaveFeedbackerData(req, resp)
	}

	return res
}

func (m *Monitor) SaveAnalyserData(req models.StdAnalyseReqBody, resp models.StdAnalyseRespBody) error {
	fmt.Println("Saved analyser data")
	return nil
}

func (m *Monitor) SaveFeedbackerData(req models.StdFeedbackReqBody, resp models.StdFeedbackRespBody) error {
	fmt.Println("Saved feedbacker data")
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
