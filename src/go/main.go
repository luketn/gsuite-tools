package main

import (
	"context"
	"encoding/base64"
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

var rootPath = LookupEnv("GSUITE_LAMBDA_ROOT_PATH", "ui")

func LookupEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	finalValue := ""
	if exists {
		finalValue = value
	} else {
		finalValue = defaultValue
	}
	return finalValue
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, evt events.APIGatewayV2HTTPRequest) (Response, error) {
	if strings.HasPrefix(evt.RawPath, "/api") {
		switch evt.RawPath {
		case "/api/pdf":
			item := []ItemDetail{
				{
					Name:   "Front End Consultation",
					Desc:   "Experience Review",
					Amount: 150,
					Qty:    4,
					Total:  600,
				},
			}
			data := Invoice{
				InvoiceNo:   "Invoice",
				InvoiceDate: "January 1, 2019",
				Currency:    "AUD",
				AmountDue:   600,
				Items:       item,
				Total:       600,
			}
			invoice, err := CreateInvoice("", data)
			if err != nil {
				println("Error generating pdf %s", err)
				return Response{StatusCode: 404}, errors.New("Failed to generate the pdf")
			} else {
				return Response{
					StatusCode: 200,
					Headers: map[string]string{
						"content-type": "application/pdf",
					},
					Body:            base64.StdEncoding.EncodeToString(invoice),
					IsBase64Encoded: true,
				}, nil
			}
		default:
			return Response{StatusCode: 404}, errors.New("Api not defined: " + evt.RawPath)
		}
	} else {
		content, contentType, err := GetStaticContent(rootPath, evt.RawPath)
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
