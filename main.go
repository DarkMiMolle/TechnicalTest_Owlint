package main

import (
	"github.com/DarkMiMolle/TechnicalTest_Owlint/server"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/util"
	"github.com/gin-gonic/gin"
	"os"
)

var IP = "localhost"

func main() {
	if res := os.Getenv("LOCALHOST"); res != "" {
		IP = res
	}
	router := gin.Default()

	// GET request
	router.GET("/target/:targetId/comments", server.GetTargetedComments)

	// POST request
	router.POST("/target/:targetId/comments", server.PostReplyComment)

	// run server, listening on port 8080. server is run on a container
	err := router.Run(IP + ":8080")
	util.PanicErr(err) // if the server can't run there is not a lot of thing we can do but print the error
}
