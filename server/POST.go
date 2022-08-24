package server

import (
	"github.com/DarkMiMolle/TechnicalTest_Owlint/datas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostReplyComment(requestInfo *gin.Context) {
	targetId := requestInfo.Param("targetId")
	var newComment datas.Comment
	err := requestInfo.BindJSON(&newComment)
	if err != nil {
		requestInfo.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if targetId != newComment.TargetId {
		requestInfo.IndentedJSON(http.StatusBadRequest, "difference between the url targetId: "+targetId+" and the body comment targetId: "+newComment.TargetId)
		return
	}
	if err := datas.RecordComment(&newComment); err != nil {
		requestInfo.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	requestInfo.IndentedJSON(http.StatusOK, newComment)
}
