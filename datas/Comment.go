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
