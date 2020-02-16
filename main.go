package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Thing struct {
	Name     string `json:"name"`
	Info     string `json:"info"`
	Category string `json:"category"`
}

type Things []Thing

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	tableName := "Things"

	fmt.Println(request.HTTPMethod)

	if request.HTTPMethod == "GET" {
		filt := expression.Name("Category").Equal(expression.Value("general"))
		proj := expression.NamesList(expression.Name("Name"), expression.Name("Info"), expression.Name("Category"))
		expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
		if err != nil {
			fmt.Println("Got error building expression:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		params := &dynamodb.ScanInput{
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			FilterExpression:          expr.Filter(),
			ProjectionExpression:      expr.Projection(),
			TableName:                 aws.String(tableName),
		}

		result, err := svc.Scan(params)
		if err != nil {
			fmt.Println("Query API call failed:")
			fmt.Println((err.Error()))
			os.Exit(1)
		}

		body, err := json.Marshal(result.Items)
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}
		APIResponse := events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 200,
			Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		}
		return APIResponse, nil
	} else if request.HTTPMethod == "POST" {
		thing := Thing{}
		err := json.Unmarshal([]byte(request.Body), &thing)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
		}

		fmt.Println(thing)

		//Returning response with AWS Lambda Proxy Response
		return events.APIGatewayProxyResponse{Body: "Success", StatusCode: 201}, nil
	}
	return events.APIGatewayProxyResponse{Body: "Nothing to do", StatusCode: 200}, nil
}
