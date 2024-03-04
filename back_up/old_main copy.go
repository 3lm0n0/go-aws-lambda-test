package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

var httpLambda *httpadapter.HandlerAdapter

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return httpLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Initialize logger
	logger := log.New(os.Stdout, "lambda: ", log.LstdFlags)

	// Register HTTP handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := jsonResponse(w, "main")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		err := jsonResponse(w, "hello")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		err := jsonResponse(w, "world")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		err := jsonResponse(w, "pong")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})

	// Create HTTP handler adapter
	httpLambda = httpadapter.New(http.DefaultServeMux)

	// Start Lambda function
	lambda.StartHandler(Handler)
}

func jsonResponse(w http.ResponseWriter, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(map[string]string{
		"message": "JSON response: " + msg + " - " + time.Now().Format(time.UnixDate),
	})
}
