package datas

// Thread is a succession of (potential) answer link to an original Comment.
// The MainComment is the default Comment, and every other comment are Replies to it.
// It also makes the link between the asked data in the REST request, but it is not represented in the DataBase (it is only build and send as a response)
type Thread struct {
	MainComment *Comment   `json:""` /// Don't know if it exists a way to indicate MainComment is the default Comment with json tags.
	Replies     []*Comment `json:"replies"`
}
