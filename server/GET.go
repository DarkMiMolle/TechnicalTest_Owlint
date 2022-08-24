package server

import (
	"fmt"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/datas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTargetedComments
//
// function handler of the http.GET request to the url:
//
// /target/:targetId/comment
func GetTargetedComments(requestInfo *gin.Context) {
	fmt.Println("GET Targeted Comment")
	targetId := requestInfo.Param("targetId")
	comments, err := datas.GetCommentsOf(targetId)
	if err != nil {
		requestInfo.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	var threads []datas.Thread
	for _, comment := range comments {
		threads = append(threads, comment.AsThread())
	}
	fmt.Println(threads)
	requestInfo.IndentedJSON(http.StatusOK, threads)
}
