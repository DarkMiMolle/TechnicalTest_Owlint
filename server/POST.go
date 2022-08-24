package server

import (
	"encoding/json"
	"fmt"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/backend"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/datas"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const otherServiceURL = "https://faulty-backend.herokuapp.com/on_comment"

var otherService = backend.Client{
	Url: otherServiceURL,

	RetryPolicy: backend.MakeRetryPolicy(
		backend.RetryPolicyStep{
			OnStatusCode: []int{429},
			NbOfRetry:    3,
			TimeInterval: 3 * time.Second,
		},
		backend.RetryPolicyStep{
			OnStatusCode: []int{400, -500},
			NbOfRetry:    3,
			TimeInterval: 2 * time.Second,
		},
		backend.RetryPolicyStep{
			OnStatusCode: []int{500, 1000},
			NbOfRetry:    1,
			TimeInterval: 4 * time.Second,
		}),
}

func forwardToOtherService(comment datas.Comment) {
	reqBody, err := json.Marshal(map[string]string{
		"message": comment.TextFr.String(),
		"author":  comment.AuthorId,
	})
	util.PanicErr(err)

	comingResp, _ := otherService.Post("application/json", reqBody)

	receivedData := make([]byte, 2048)
	n, _ := comingResp.Get().Body.Read(receivedData)
	fmt.Println("STATUS: " + comingResp.Get().Status + "\nContent: " + string(receivedData[:n]))
}

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
	go forwardToOtherService(newComment)
	if err := datas.RecordComment(&newComment); err != nil {
		requestInfo.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	requestInfo.IndentedJSON(http.StatusOK, newComment)
}
