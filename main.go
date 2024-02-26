package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)
var httpLambda *httpadapter.HandlerAdapter

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("==> Handler: ", req.RequestContext.RequestID)
    fmt.Println("==> Handler: ", req.Path)
    fmt.Println("==> Handler: ", req.HTTPMethod)
    return httpLambda.ProxyWithContext(ctx, req)
}

func main() {
    fmt.Println("==> main: ", time.Now().Format(time.UnixDate))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /: ", r)
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - main")
	})
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /hello: ", r)
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - hello")
	})
    http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /world: ", r)
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - world")
	})
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /ping: ", r)
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - pong")
	})

	httpLambda = httpadapter.New(http.DefaultServeMux)

	lambda.Start(Handler)
}