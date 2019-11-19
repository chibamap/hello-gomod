package main

import (

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/hogehoge-banana/hello-gomod/hello"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body, err := hello.Hoge()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
