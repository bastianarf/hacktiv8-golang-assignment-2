package router

import (
	"fmt"
	"golang-hacktiv8-assignment-2/config"
	"golang-hacktiv8-assignment-2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(c controller.Controller) error {
	port := config.GetServerPortEnv()
	serverInfo := fmt.Sprintf("localhost:%s", port)

	r := gin.Default()

	r.POST("/orders", c.CreateOrder)
	r.GET("/orders", c.GetOrders)
	r.PUT("/orders/:orderId", c.UpdateOrder)
	r.DELETE("/orders/:orderId", c.DeleteOrder)

	return r.Run(serverInfo)
}
