package models

type StdAnalyseReqBody struct{}
type StdAnalyseRespBody struct {
	Error error
}

type StdFeedbackReqBody struct{}
type StdFeedbackRespBody struct {
	Error error
}

type Feedbacker interface {
	Feedback(StdFeedbackReqBody) StdFeedbackRespBody
}

type Analyser interface {
	Analyse(StdAnalyseReqBody) StdAnalyseRespBody
}
