package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func handler(ctx context.Context) (map[string]interface{}, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("resume-visitor-counter-terraform"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("visitor-count"),
			},
		},
		UpdateExpression: aws.String("SET counterNumber = counterNumber + :inc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":inc": {
				// increment by 1
				N: aws.String("1"),
			},
		},
		// return the updated attributes
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	result, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Failed to update item: %v", err)
	}

	//return map[string]interface{}{"count": newCount}, nil
	return map[string]interface{}{"count": *result.Attributes["counterNumber"].N}, nil
}

func main() {

	// uncomment this after local test
	lambda.Start(handler)

	/*
		//for local test; left here for history :)
		res, err := handler(context.Background())
		if err != nil {
			log.Fatalf("handler error: %v", err)
		}
		fmt.Println(res)
	*/

}
