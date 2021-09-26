package main

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var rootPath, rootPathSet = os.LookupEnv("GSUITE_LAMBDA_ROOT_PATH")

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, evt events.APIGatewayV2HTTPRequest) (Response, error) {
	if strings.HasPrefix(evt.RawPath, "/api") {
		switch evt.RawPath {
		default:
			return Response{StatusCode: 404}, errors.New("Api not defined: " + evt.RawPath)
		}
	} else {
		finalRootPath := ""
		if rootPathSet {
			finalRootPath = rootPath
		} else {
			finalRootPath = "ui"
		}
		content, contentType, err := GetStaticContent(finalRootPath, evt.RawPath)
		if err != nil {
			return Response{StatusCode: 404}, err
		} else {
			return Response{
				StatusCode: 200,
				Headers: map[string]string{
					"content-type": contentType,
				},
				Body: content,
			}, nil
		}
	}
}

func main() {
	lambda.Start(Handler)
}
