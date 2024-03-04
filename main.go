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
		loggerLog(r, logger, "/")
		err := jsonResponse(w, "welcome to main :)")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		loggerLog(r, logger, "/hello world")
		err := jsonResponse(w, "hello world")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		loggerLog(r, logger, "/world, hello")
		err := jsonResponse(w, "world, hello")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		loggerLog(r, logger, "/ping")
		err := jsonResponse(w, "pong")
		if err != nil {
			logger.Printf("Error encoding JSON response: %v", err)
		}
	})

	httpLambda = httpadapter.New(http.DefaultServeMux)

	lambda.Start(Handler)
}

func jsonResponse(w http.ResponseWriter, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(map[string]string{
		"message": "JSON response: " + msg + " - " + time.Now().Format(time.UnixDate),
	})
}

func loggerLog(r *http.Request, logger *log.Logger, path string) {
	logger.Println(
		"\n• HandleFunc: ",                             path,
		"\n• Request: ",                                r,
		"\n• r.Header.Get(routeKey): ",                 r.Header.Get("routeKey"),
		"\n• r.Header.Get(Path): ",                     r.Header.Get("Path"),
		"\n• URL: ",                                   	r.URL,
		"\n• Path: ",                                   r.URL.Path, 
		"\n• Method: ",                                 r.Method,
		"\n• Address: ",                                r.RemoteAddr, 
		"\n• Agent: ",                                  r.UserAgent(), 
		"\n• Content-Type: ",                           r.Header.Get("Content-Type"), 
		"\n• Origin: ",                                 r.Header.Get("Origin"),
		"\n• Access-Control-Allow-Origin: ",            r.Header.Get("Access-Control-Allow-Origin"),
		"\n• Access-Control-Allow-Methods: ",           r.Header.Get("Access-Control-Allow-Methods"),
		"\n• Access-Control-Allow-Headers: ",           r.Header.Get("Access-Control-Allow-Headers"),
		"\n• Access-Control-Request-Method: ",          r.Header.Get("Access-Control-Request-Method"),
		"\n• Access-Control-Request-Headers: ",         r.Header.Get("Access-Control-Request-Headers"),
		"\n• Access-Control-Request-Expose-Headers: ",  r.Header.Get("Access-Control-Request-Expose-Headers"),
		"\n• Access-Control-Request-Credentials: ",     r.Header.Get("Access-Control-Request-Credentials"),
	)
}