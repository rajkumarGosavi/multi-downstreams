package providera

import (
	"abstraction/models"
	"fmt"
)

type ProviderA struct {
	key string
}

func New() *ProviderA {
	return &ProviderA{
		key: "qwerty1",
	}
}

func (a *ProviderA) Analyse(models.StdAnalyseReqBody) models.StdAnalyseRespBody {
	fmt.Println("PAApi1")
	return models.StdAnalyseRespBody{OP: "PA Analyse"}
}

func (a *ProviderA) Feedback(models.StdFeedbackReqBody) models.StdFeedbackRespBody {
	fmt.Println("PAApi2")
	return models.StdFeedbackRespBody{OP: "PA Feedback"}
}
