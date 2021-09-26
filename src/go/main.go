package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var index []byte = nil
var err error = nil

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, evt events.APIGatewayV2HTTPRequest) (Response, error) {
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(index),
		Headers: map[string]string{
			"Content-Type":           "text/html",
			"X-MyCompany-Func-Reply": "world-handler",
			"requestedPath":          evt.RawPath,
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
