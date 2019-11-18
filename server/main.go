package main

import (
	"fmt"
	"log"
	"net/http"

	"./config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rs/cors"
)

var sess, _ = session.NewSession(&aws.Config{
	Region: aws.String("us-east-1"),
})

func main() {
	// rds.Start()
	router := serverRouter()

	fmt.Println(`server v`, config.Version, ` listening on port :8080`)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
