package datas

import (
	"context"
	"fmt"
	. "github.com/DarkMiMolle/TechnicalTest_Owlint/database"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/util"
	translate "github.com/bas24/googletranslatefree"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

// Text can be translate from en <-> fr
type Text string

// TranslateInFrEn translate the text in both languages fr and en
func (text Text) TranslateInFrEn() (fr, en Text, err error) {
	if text == "" {
		return "", "", nil
	}
	frStr, err := translate.Translate(string(text), "fr", "en")
	fr = Text(frStr)
	if err == nil {
		return
	}
	enStr, err := translate.Translate(string(text), "en", "fr")
	en = Text(enStr)
	return
}

// TranslateToEn translate to english
func (text Text) TranslateToEn() (en Text, err error) {
	if text == "" {
		return "", nil
	}
	enStr, err := translate.Translate(string(text), "fr", "en")
	en = Text(enStr)
	return
}

// TranslateToFr translate to french
func (text Text) TranslateToFr() (fr Text, err error) {
	if text == "" {
		return "", nil
	}
	frStr, err := translate.Translate(string(text), "en", "fr")
	fr = Text(frStr)
	return
}
func (text Text) String() string { return string(text) }

// timestamp allows to use time.Time but will be json-encoded as asked by the API requirements.
type timestamp struct {
	time.Time
}

func (t timestamp) MarshalJSON() ([]byte, error) {
	str := fmt.Sprint(t.Unix())
	return []byte(strconv.Quote(str)), nil
}
func (t *timestamp) UnmarshalJSON(content []byte) error {
	str := string(content)
	str, err := strconv.Unquote(str)
	if err != nil {
		return err
	}
	timestamp, err := strconv.ParseInt(str, 10, 64)
	t.Time = time.Unix(timestamp, 0)
	return err
}

// Comment is the representation of a comment in the DataBase (and through the API argument it is a Comment without replies)
type Comment struct {
	Id          string    `json:"id,omitempty"`
	TextFr      Text      `json:"textFr,omitempty"`
	TextEn      Text      `json:"textEn,omitempty"`
	PublishedAt timestamp `json:"publishedAt,omitempty"` // PublishedAt is a string in the json file
	AuthorId    string    `json:"authorId,omitempty"`
	TargetId    string    `json:"targetId,omitempty"`
}

// lastComment to remove ? depreciated
var lastComment = &Comment{}

// GetComment to remove ? depreciated
func GetComment(id string) *Comment {
	if lastComment.Id != id {
		res := DataBase().FindOne(context.Background(), bson.D{{"id", id}})
		if res == nil {
			fmt.Println(id, "Not presents.")
			return nil
		}
		util.PanicErr(res.Decode(lastComment))
	}
	return lastComment
}

// GetCommentsOf return all the commentary linked to the given targetId from the database
func GetCommentsOf(targetId string) ([]Comment, error) {
	sortOpt := options.Find().SetSort(bson.D{{"publishedat", 1}})
	findRes, err := DataBase().Find(context.Background(), bson.D{{"targetid", targetId}}, sortOpt)
	if err != nil {
		return nil, err
	}
	var multipleComment []Comment
	fmt.Println(multipleComment)
	err = findRes.All(context.Background(), &multipleComment)
	if err != nil {
		return nil, err
	}
	return multipleComment, nil
}

// RecordComment record (or update) a comment in the database
func RecordComment(comment Comment) error {
	res := DataBase().FindOneAndReplace(context.Background(), bson.D{{"id", comment.Id}}, comment)
	if res.Err() == nil { // Done
		return nil
	}
	_, err := DataBase().InsertOne(context.Background(), comment)
	return err
}
