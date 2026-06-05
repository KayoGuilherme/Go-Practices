package main

import (
	"log"
	"os"

	"go.uber.org/zap"

	"server/database"
	"server/internal/handler"
	"server/internal/infrastructure/repository"
	wsinfra "server/internal/infrastructure/websocket"
	"server/internal/service"
	"server/router"
)

func main() {
	dbConnection, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	var logger *zap.Logger

    if os.Getenv("APP_ENV") == "production" {
        logger, _ = zap.NewProduction()
    } else {
        logger, _ = zap.NewDevelopment()
    }
    defer logger.Sync()

	// Users
	userRepo := repository.NewRepository(dbConnection.GetDB())
	userSvc := service.NewService(userRepo)
	userHandler := handler.NewUserHandler(userSvc, logger)

	// Auth
	authRepo := repository.NewRepository(dbConnection.GetDB())
	authSvc := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authSvc)

	// product 
	productRepo := repository.NewProductRepository(dbConnection.GetDB())
	productsvc := service.NewProductService(productRepo);
	productHandler := handler.NewProductHandler(productsvc)

	// Category
	categoryrepo := repository.NewCategoryRepository(dbConnection.GetDB())
	categorysvc := service.NewCategoryService(categoryrepo)
	categoryHanddler := handler.NewCategoryHandler(categorysvc)

	// Product ↔ Category (many-to-many)
	linkRepo := repository.NewCategoryProductRepository(dbConnection.GetDB())
	linkSvc := service.NewCategoryProductService(linkRepo, productRepo, categoryrepo)
	linkHandler := handler.NewCategoryProductHandler(linkSvc)

	// WS infra 
	hub := wsinfra.NewHub()
	wsHandler := handler.NewWSHandler(hub)
	go wsinfra.Run(hub)

	router.InitRouter(userHandler, authHandler, productHandler, categoryHanddler, linkHandler, wsHandler)
	router.Start("0.0.0.0:8081")
}
