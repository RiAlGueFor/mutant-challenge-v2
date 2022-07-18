package mutant

import (
  "encoding/json"
  "github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{})(*events.APIGatewayProxyResponse, error){
  resp:=events.APIGatewayProxyResponse(Headers: map[string]string["Content-Type:":"application/json"])
  resp.StatusCode = StatusCode

  stringBody, _ := json.Marshal(body)

  resp.Body = string(stringBody)
  return &resp,nil
}
