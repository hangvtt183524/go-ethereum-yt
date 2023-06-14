package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hangvtt183524/go-serverless-yt/pkg/handlers"
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	dynaClient := dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyRequest, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)

	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)

	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)

	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)

	default:
		return handlers.UnhandleMethod()
	}
}
