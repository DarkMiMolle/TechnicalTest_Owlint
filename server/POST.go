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
	if newComment.TextEn == "" {
		en, err := newComment.TextFr.TranslateToEn()
		if err != nil {
			newComment.TextEn = "Unable to translate that sentence"
		} else {
			newComment.TextEn = en
		}
	} else if newComment.TextFr == "" {
		fr, err := newComment.TextEn.TranslateToFr()
		if err != nil {
			newComment.TextFr = "Impossible de traduire cette phrase"
		} else {
			newComment.TextFr = fr
		}
	}
	if err := datas.RecordComment(&newComment); err != nil {
		requestInfo.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	requestInfo.IndentedJSON(http.StatusOK, newComment)
}
