package datas

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

type comment_test_Sample struct {
	comment  Comment
	asString string
}

var (
	commentObi = comment_test_Sample{
		Comment{
			Id:          "0",
			TextFr:      "Salut toi !",
			TextEn:      "Hello there !",
			PublishedAt: timestamp{time.Unix(1802310899, 0)},
			AuthorId:    "Obi-Wan",
			TargetId:    "Grievous",
		},
		`{"id":"0","textFr":"Salut toi !","textEn":"Hello there !","publishedAt":"1802310899","authorId":"Obi-Wan","targetId":"Grievous"}`,
	}
	commentGrievous = comment_test_Sample{
		Comment{
			Id:          "1",
			TextFr:      "General Kenobi !",
			TextEn:      "General Kenobi !",
			PublishedAt: timestamp{time.Unix(1639477064, 0)},
			AuthorId:    "Grievous",
			TargetId:    "Obi-Wan",
		},
		`{"id":"1","textFr":"General Kenobi !","textEn":"General Kenobi !","publishedAt":"1639477064","authorId":"Grievous","targetId":"Obi-Wan"}`,
	}
	commentHello = comment_test_Sample{
		Comment{Id: "0", TextFr: "Salut", TextEn: "Hi", PublishedAt: timestamp{time.Unix(1802310899, 0)}, AuthorId: "Flo", TargetId: "TestComment"},
		`{"id":"0","textFr":"Salut","textEn":"Hi","publishedAt":"1802310899","authorId":"Flo","targetId":"TestComment"}`,
	}
)

func TestComment_JsonUnmarshal(t1 *testing.T) {
	type fields = Comment
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Add test cases.
		{
			"ValidComment",
			commentObi.comment,
			args{
				[]byte(commentObi.asString),
			},
			false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			var result Comment

			if err := json.Unmarshal(tt.args.content, &result); (err != nil) != tt.wantErr {
				t1.Errorf("json.Unmarshal(Comment) error = %v, wantErr %v", err, tt.wantErr)
			} else if !reflect.DeepEqual(result, tt.fields) {
				t1.Errorf("Data is not the expected one.\ngot:\n\t%v\nwant:\n\t%v", result, tt.fields)
			}
		})
	}
}

func TestComment_JsonMarshal(t1 *testing.T) {
	type fields = Comment
	tests := []struct {
		name    string
		fields  Comment
		want    string
		wantErr bool
	}{
		// Add test cases.
		{
			"Valid comment",
			commentGrievous.comment,
			commentGrievous.asString,
			false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tt.fields
			got, err := json.Marshal(t)
			if (err != nil) != tt.wantErr {
				t1.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t1.Errorf("MarshalJSON() got\n\t%v\nwant\n\t%v", string(got), tt.want)
			}
		})
	}
}
