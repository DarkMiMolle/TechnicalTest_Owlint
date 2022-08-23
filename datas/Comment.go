package datas

import "time"

// Text can be translate from en <-> fr
type Text string

// timestamp allows to use time.Time but will be json-encoded as asked by the API requirements.
type timestamp struct {
	time.Time
}

// Comment is the representation of a comment in the DataBase (and through the API argument)
type Comment struct {
	Id          string    `json:"id,omitempty"`
	TextFr      Text      `json:"textFr"`
	TextEn      Text      `json:"textEn"`
	PublishedAt timestamp `json:"publishedAt"` // PublishedAt is a string in the json file
	AuthorId    string    `json:"authorId"`
	TargetId    string    `json:"targetId"`
}

// Thread is a succession of (potential) answer link to an original Comment.
// The MainComment is the default Comment, and every other comment are Replies to it.
// It also makes the link between the asked data in the REST request, but it is not represented in the DataBase (it is only build and send as a response)
type Thread struct {
	MainComment *Comment   `json:""` /// Don't know if it exists a way to indicate MainComment is the default Comment with json tags.
	Replies     []*Comment `json:"replies"`
}
