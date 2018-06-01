package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

// Request ...
type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

// Response ..
type Response struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

// Handler ...
func Handler(request Request) (Response, error){
	return Response{
		Message: fmt.Sprintf("Processed request %f", request.ID),
		OK:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
