# Introduction

Given we have some downstream providers that provides us some Analytics service,
for that they need necessary data, every provider will have a unique request data 
and will give us some response in return.

Once we get the analysis we then review it for our system and then we share a feedback with
the providers for the given analysis.


This module gives us an example of how we can achieve the above use case.

## Code Structure:

```
├── go.mod
├── main.go
├── models
│   └── model.go
├── monitor
│   └── monitor.go
├── providera
│   └── a.go
├── providerb
│   └── b.go
└── README.md
```

`main.go` -> The entry point which will iterate over the events, that we need to send either for analysis or as a feedback

`models/model.go` -> contains our requests and responses types and interface defination

`monitor/monitor.go` -> is going to overlook the entire execution of our operations, it will have various resource access like DB etc.

`provider<x>/` -> has the actual implementations for the downstreams

## The Flow

### Monitor package

```go
type Monitor struct {
	analysers   []models.Analyser
	feedbackers []models.Feedbacker
	app         interface{}
}
```

The `Monitor` will over look all the activities that needs to be taken care in the service. Only it has the access to the resources such as DB, redis, File Systems, Cloud Resources etc. referred by the `app` field

It has the `analysers` field that has all the providers that gives us the capability to get our data analysed by integrating with them,
and following that we have `feedbackers` field which has the providers which allow us to give feedback on their analysis response.

To be a member of the the `analysers` slice any concrete type has to satisfy `Analyser` interface
```go
type Analyser interface {
	Analyse(StdAnalyseReqBody) StdAnalyseRespBody
}
```

Similarly a type has to satisfy the `Feedbacker` interface to be part of `feedbackers` slice.

```go
type Feedbacker interface {
	Feedback(StdFeedbackReqBody) StdFeedbackRespBody
}
```

This will help us abstract the concrete provider specific implementation details.

Both the `Analyse(StdAnalyseReqBody) StdAnalyseRespBody` and `Feedback(StdFeedbackReqBody) StdFeedbackRespBody` methods take a standard/general input and return a general output. Thus we can send fields as and when necessary for the other providers without breaking the contract with the already existing providers.


The `Monitor` has methods `Analyse(event interface{})` and `Feedback(event interface{})`
that respectively takes upon the task to generate the necessary standard request for invoking all the analysers and feedbackers available.


### Providera package

The actual implementation of the provider A's APIs and necessary methods required to make the communication successful.

### Providerb package

The actual implementation of the provider B's APIs and necessary methods required to make the communication successful.


## Testing

Since all the provider specific implementation is abstracted our monitor can mock it can we can test it easily. 