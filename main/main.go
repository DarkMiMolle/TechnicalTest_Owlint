package main

import (
	"fmt"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/util"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// GET request
	router.GET("/target/:targetId/comments", func(*gin.Context) { fmt.Println("TODO: get targetId comment") })

	// POST request
	router.POST("/", func(*gin.Context) { fmt.Println("TODO: post new target and POST it in another back-end") })

	// run server, listening on port 8080. server is run on a container
	err := router.Run("localhost:8080")
	util.PanicErr(err) // if the server can't run there is not a lot of thing we can do but print the error
}
