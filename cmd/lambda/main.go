package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"infantry/bindings"
)

func HandleRequest(ctx context.Context, name bindings.Plan) (bindings.Report, error) {
	var report bindings.Report
	return report, nil
}

func main() {
	lambda.Start(HandleRequest)
}
