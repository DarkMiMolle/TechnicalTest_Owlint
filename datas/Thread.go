package datas

import (
	"encoding/json"
	"strings"
)

// Thread is a succession of (potential) answer link to an original Comment.
// The MainComment is the default Comment, and every other comment are Replies to it.
// It also makes the link between the asked data in the REST request, but it is not represented in the DataBase (it is only build and send as a response)
type Thread struct {
	MainComment *Comment   `json:""` /// Don't know if it exists a way to indicate MainComment is the default Comment with json tags.
	Replies     []*Comment `json:"replies"`
}

func (thread Thread) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal(*thread.MainComment)
	if err != nil {
		return ret, err
	}
	str := string(ret[:len(ret)-1])
	str += ",\"replies\":"
	ret, err = json.Marshal(thread.Replies)
	str += string(ret) + "}"
	return []byte(str), err
}

func (thread *Thread) UnmarshalJSON(jsonStr []byte) error {
	if thread.MainComment == nil {
		thread.MainComment = new(Comment)
	}
	err := json.Unmarshal(jsonStr, thread.MainComment)
	if err != nil {
		return err
	}
	if thread.Replies == nil {
		thread.Replies = []*Comment{}
	}
	leftStr := string(jsonStr)[strings.Index(string(jsonStr), "["):]
	leftStr = leftStr[:len(leftStr)-1]
	err = json.Unmarshal([]byte(leftStr), &thread.Replies)
	return err
}

func (thread Thread) String() string {
	res, err := thread.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(res), ",", ", ")
}

func GetThreadOf(comment *Comment) Thread {
	thread := Thread{MainComment: comment}
	thread.Replies, _ = GetCommentsOf(comment.Id)
	return thread
}
