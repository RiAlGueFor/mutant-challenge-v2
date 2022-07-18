package main

import(
  "os"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
  "github.com/RiAlGueFor/mutant-challenge-v2/internal/mutant"
)

var (
  dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
  region:=os.Getenv("AWS_REGION")
  awsSession, err:=session.NewSession(&aws.Config{
    Region: aws.String(region)},)

  if err!=nil{
    return
  }

  dynaClient = dynamodb.New(awsSession)
  lambda.Start(handler)

}

const tableName = "LambaDNAValidationRecords"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error){
    switch req.HTTPMethod {
      case "GET":
        return mutant.GetStats(req,tableName,dynaClient)
      case "POST":
        return mutant.CheckMutantDNA(req,tableName,dynaClient)
      default:
        return mutant.UnhandledMethod()
    }
}
