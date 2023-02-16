package providerb

import (
	"abstraction/models"
	"fmt"
)

type ProviderB struct {
	key string
}

func New() *ProviderB {
	return &ProviderB{
		key: "qwerty2",
	}
}

func (b *ProviderB) Analyse(models.StdAnalyseReqBody) models.StdAnalyseRespBody {
	fmt.Println("PBApi1")
	return models.StdAnalyseRespBody{}
}

func (b *ProviderB) Feedback(models.StdFeedbackReqBody) models.StdFeedbackRespBody {
	fmt.Println("PBApi2")
	return models.StdFeedbackRespBody{}
}
