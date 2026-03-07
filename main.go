package main

import (
	"context"
	"log"
	"transmitter/push"

	qService "transmitter/sqs"

	tservice "transmitter/target_service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda
var sqsClient *sqs.Client

func init() {
	// 1. Initialize Gin Router
	r := gin.Default()

	// 2. Initialize AWS Config & SQS Client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	sqsClient = sqs.NewFromConfig(cfg)

	qService := qService.NewQueue(&cfg)
	tservice := tservice.NewTargetService()
	pService := push.NewPushService(qService, tservice)
	pushHandler := push.NewPushHandler(pService)

	// 3. Define the Route
	r.POST("/push", pushHandler.HandlePush)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	// 4. Wrap Gin with the Lambda Adapter
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request, default to "World"
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
