package main

import (
	"os"

	"data-pusher/connectors"
	"data-pusher/controller"
	"data-pusher/repository"

	"github.com/labstack/echo"
)

type Container struct {
	AccountInstance     controller.AccountController
	DestinationInstance controller.DestinationController
	DataHandler         controller.DataHandlerController
}

func LoadContainer() *Container {
	return &Container{
		AccountInstance:     controller.AccountController{Mysql: repository.SingletonMysqlCon()},
		DestinationInstance: controller.DestinationController{Mysql: repository.SingletonMysqlCon()},
		DataHandler:         controller.DataHandlerController{Mysql: repository.SingletonMysqlCon()},
	}
}
func init() {
	connectors.LoadEnv()
}

func main() {
	containerInstance := LoadContainer()
	e := echo.New()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	// Accounts Routes
	e.POST("/account", containerInstance.AccountInstance.CreateAccount)
	e.PUT("/account", containerInstance.AccountInstance.UpdateAccount)
	e.GET("/account/", containerInstance.AccountInstance.GetAccountDetails)
	e.DELETE("/account", containerInstance.AccountInstance.DeleteAccount)

	// Destination Routes

	e.POST("/destination", containerInstance.DestinationInstance.CreateDestination)
	e.GET("/destination/", containerInstance.DestinationInstance.GetDestinationDetails)

	// Data handler
	e.GET("/incoming_data", containerInstance.DataHandler.HandleData)

	e.Start(":" + PORT)
}
