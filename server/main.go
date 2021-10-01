package main

import (
	"os"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// These are the endpoint
	// Create
	router.POST("/order/create", routes.AddOrder)
	// Read
	router.GET("/waiter/:waiter", routes.GetOrdersByWaiter)
	router.GET("/orders", routes.GetOrders)
	router.GET("/order/:id", routes.GetOrderById)
	//Update
	router.PUT("/waiter/update/:id", routes.UpdateWaiter)
	router.PUT("/order/update/:id", routes.UpdateOrder)
	//Delete
	router.DELETE("/order/delete/:id", routes.DeleteOrder)

	//This runs the server and allows it to listen to request
	router.Run(":" + port)
}
