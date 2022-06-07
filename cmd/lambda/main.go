package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"infantry/bindings"
	"infantry/engine"
)

func HandleRequest(ctx context.Context, plan bindings.Plan) (bindings.Report, error) {
	var report, err = engine.Start(plan)
	// TODO: Log stuff n junk yo
	return report, err
}

func main() {
	lambda.Start(HandleRequest)
}
