package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nddat1811/golang_mongodb/controllers"
	"github.com/nddat1811/golang_mongodb/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	docs "github.com/nddat1811/golang_mongodb/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server      *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("userdb").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	server = gin.Default()
}

// @title Swagger demo MongoDB service API
// @version 1.0
// @description This is demo server
// @host localhost:9090
// @BasePath /v1
// @securityDeifinitions.basic BasicAuth
// @securityDefinitions.apiKey ApiKeyAuth
// @in hearder
// @name Authorization

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	docs.SwaggerInfo.BasePath = "/v1"
	uc.RegisterUserRoutes(basepath)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(":9090"))
}
