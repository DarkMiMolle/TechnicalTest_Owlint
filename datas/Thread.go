package datas

import (
	"encoding/json"
	"strings"
)

// A Thread is Comment with all replies
// Each Replies is also a Thread (or a Comment if Replies is nil)
// It also makes the link between the asked data in the REST request, but it is not represented in the DataBase (it is only build and send as a response)
type Thread struct {
	Comment
	Replies []Thread `json:"replies"`
}

func (thread Thread) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal(&thread.Comment)
	if err != nil {
		return ret, err
	}
	if thread.Replies == nil {
		return ret, nil
	}
	str := string(ret[:len(ret)-1])
	str += ",\"replies\":"
	ret, err = json.Marshal(thread.Replies)
	str += string(ret) + "}"
	return []byte(str), err
}
func (thread *Thread) UnmarshalJSON(jsonStr []byte) error {
	err := json.Unmarshal(jsonStr, &thread)
	return err
}
func (thread Thread) String() string {
	res, err := thread.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(res), ",", ", ")
}

// AsThread transform a simple Comment to a full Thread
func (comment Comment) AsThread() Thread {
	thread := Thread{Comment: comment}
	if comment.Id == comment.TargetId { // Can't make a Thread with itself
		return thread
	}
	comments, _ := GetCommentsOf(comment.Id)
	for _, comment := range comments {
		thread.Replies = append(thread.Replies, comment.AsThread())
	}
	return thread
}

// AsThread overriding the Comment implem of that method since it is already a Thread.
func (th Thread) AsThread() Thread {
	return th
}
