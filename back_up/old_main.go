package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)
var httpLambda *httpadapter.HandlerAdapter

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("==> Handler req.RequestContext.RequestID: ", req.RequestContext.RequestID)
    fmt.Println("==> Handler req.Path: ", req.Path)
    fmt.Println("==> Handler req.HTTPMethod: ", req.HTTPMethod)
    fmt.Println("==> Handler req: ", req)
    log.Print(
			"\n==> Logs: ",
			"\n• Path: ",                                   req.URL.Path, 
			"\n• Method: ",                                 req.Method,
			"\n• Address: ",                                req.RemoteAddr, 
			"\n• Agent: ",                                  req.UserAgent(), 
			"\n• Content-Type: ",                           req.Header.Get("Content-Type"), 
			"\n• Origin: ",                                 req.Header.Get("Origin"),
			"\n• Access-Control-Allow-Origin: ",            req.Header.Get("Access-Control-Allow-Origin"),
			"\n• Access-Control-Allow-Methods: ",           req.Header.Get("Access-Control-Allow-Methods"),
			"\n• Access-Control-Allow-Headers: ",           req.Header.Get("Access-Control-Allow-Headers"),
			"\n• Access-Control-Request-Method: ",          req.Header.Get("Access-Control-Request-Method"),
			"\n• Access-Control-Request-Headers: ",         req.Header.Get("Access-Control-Request-Headers"),
			"\n• Access-Control-Request-Expose-Headers: ",  req.Header.Get("Access-Control-Request-Expose-Headers"),
			"\n• Access-Control-Request-Credentials: ",     req.Header.Get("Access-Control-Request-Credentials"),
		)
    return httpLambda.ProxyWithContext(ctx, req)
}

func old_main() {
    fmt.Println("==> main: ", time.Now().Format(time.UnixDate))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /: ", r)
        fmt.Println("==> r.Header.Get(Path): ", r.Header.Get("Path"))
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - main")
	})
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /hello: ", r)
        fmt.Println("==> r.Header.Get(Path): ", r.Header.Get("Path"))
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - hello")
	})
    http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /world: ", r)
        fmt.Println("==> r.Header.Get(Path): ", r.Header.Get("Path"))
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - world")
	})
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("==> /ping: ", r)
        fmt.Println("==> r.Header.Get(Path): ", r.Header.Get("Path"))
        fmt.Println("==> request url path: ", r.URL.Path)
        fmt.Println("==> request Header: ", r.Header)
        fmt.Println("==> request Method: ", r.Method)
        fmt.Println("==> request RequestURI: ", r.RequestURI)
		json.NewEncoder(w).Encode("JSON response: " + time.Now().Format(time.UnixDate) + " - pong")
	})

	httpLambda = httpadapter.New(http.DefaultServeMux)

	lambda.Start(Handler)
}