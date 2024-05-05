package main

import (
	"fmt"
	"log"
	"net/http"
	cfg "rub_buddy/configs"
	"rub_buddy/helper"
	"rub_buddy/routes"
	"rub_buddy/utils/bucket"
	"rub_buddy/utils/database"

	dataUser "rub_buddy/features/users/data"
	handlerUser "rub_buddy/features/users/handler"
	serviceUser "rub_buddy/features/users/service"

	dataCollector "rub_buddy/features/collectors/data"
	handlerCollector "rub_buddy/features/collectors/handler"
	serviceCollector "rub_buddy/features/collectors/service"

	dataPickup "rub_buddy/features/pickup_request/data"
	handlerPickup "rub_buddy/features/pickup_request/handler"
	servicePickup "rub_buddy/features/pickup_request/service"

	dataPickupTransaction "rub_buddy/features/pickup_transaction/data"
	handlerPickupTransaction "rub_buddy/features/pickup_transaction/handler"
	servicePickupTransaction "rub_buddy/features/pickup_transaction/service"

	"github.com/labstack/echo/v4"
)

func main() {
	var config = cfg.InitConfig()
	db, err := database.InitDB(*config)

	database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	jwtInterface := helper.New(config.Secret)

	bucketInterface, err := bucket.New(config.ProjectID, config.BucketName)
	if err != nil {
		log.Fatal(err)
	}

	userModel := dataUser.New(db)
	userService := serviceUser.New(userModel, jwtInterface)
	userController := handlerUser.NewHandler(userService, jwtInterface)

	collectorModel := dataCollector.New(db)
	collectorService := serviceCollector.New(collectorModel, jwtInterface)
	collectorController := handlerCollector.NewHandler(collectorService, jwtInterface)

	pickupModel := dataPickup.New(db)
	pickupService := servicePickup.New(pickupModel)
	pickupController := handlerPickup.NewHandler(pickupService, jwtInterface)

	pickupTransactionModel := dataPickupTransaction.New(db)
	pickupTransactionService := servicePickupTransaction.New(pickupTransactionModel)
	pickupTransactionController := handlerPickupTransaction.NewHandler(pickupTransactionService, jwtInterface)

	routes.RouteUser(e, userController, *config)
	routes.RouteCollector(e, collectorController, *config)
	routes.RoutePickup(e, pickupController, *config)
	routes.RouteTransaction(e, pickupTransactionController, *config)
	routes.RouteMedia(e, bucketInterface)
	e.Logger.Debug(db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8080)).Error())
}
