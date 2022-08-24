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

func (text Text) TranslateInFrEn() (fr, en string, err error) {
	fr, err = translate.Translate(string(text), "fr", "en")
	if err == nil {
		return
	}
	en, err = translate.Translate(string(text), "en", "fr")
	return
}
func (text Text) TranslateToEn() (en string, err error) {
	return translate.Translate(string(text), "fr", "en")
}
func (text Text) TranslateToFr() (fr string, err error) {
	return translate.Translate(string(text), "en", "fr")
}

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

// Comment is the representation of a comment in the DataBase (and through the API argument)
type Comment struct {
	Id          string    `json:"id,omitempty"`
	TextFr      Text      `json:"textFr"`
	TextEn      Text      `json:"textEn"`
	PublishedAt timestamp `json:"publishedAt"` // PublishedAt is a string in the json file
	AuthorId    string    `json:"authorId"`
	TargetId    string    `json:"targetId"`
}

var lastComment = &Comment{}

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
func GetCommentsOf(targetId string) ([]*Comment, error) {
	sortOpt := options.Find().SetSort(bson.D{{"publishedat", 1}})
	findRes, err := DataBase().Find(context.Background(), bson.D{{"targetid", targetId}}, sortOpt)
	if err != nil {
		return nil, err
	}
	var multipleComment []*Comment
	fmt.Println(multipleComment)
	err = findRes.All(context.Background(), &multipleComment)
	if err != nil {
		return nil, err
	}
	return multipleComment, nil
}
func RecordComment(comment *Comment) {
	DataBase().InsertOne(context.Background(), comment)
}
