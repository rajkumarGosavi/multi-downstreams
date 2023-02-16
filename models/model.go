package models

type StdAnalyseReqBody struct{}
type StdAnalyseRespBody struct {
	OP    string
	Error error
}

type StdFeedbackReqBody struct{}
type StdFeedbackRespBody struct {
	OP    string
	Error error
}

type Feedbacker interface {
	Feedback(StdFeedbackReqBody) StdFeedbackRespBody
}

type Analyser interface {
	Analyse(StdAnalyseReqBody) StdAnalyseRespBody
}

type AnalyseChanData struct {
	Req  StdAnalyseReqBody
	Resp StdAnalyseRespBody
}

type FeedbackChanData struct {
	Req  StdFeedbackReqBody
	Resp StdFeedbackRespBody
}
