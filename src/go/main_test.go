package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"io/ioutil"
	"testing"
)

func TestHandler(t *testing.T) {
	rootPath = "../resources/ui"

	ctx := context.Background()
	inputJSON, err := ioutil.ReadFile("../testData/api-request.json")
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	// de-serialize into Go object
	var inputEvent events.APIGatewayV2HTTPRequest
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}
	response, err := Handler(ctx, inputEvent)
	if err != nil {
		t.Errorf("Handler failed: %v", err)
	} else {
		t.Logf("Success!\n%v", response)
	}
}
