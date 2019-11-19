package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/hogehoge-banana/hello-gomod/fuga"
	"github.com/hogehoge-banana/hello-gomod/hello"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var body string
	var err error
	switch request.HTTPMethod {
	case "POST":
		body, err = hello.Hoge()
	case "GET":
		body, err = fuga.Welcome()
	default:

	}

	if err != nil {
		body = fmt.Sprintf("Hello, %v", err)
		return events.APIGatewayProxyResponse{
			Body:       body,
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
