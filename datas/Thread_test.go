package datas

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func (c comment_test_Sample) AsMainCommentStr(replies []comment_test_Sample) string {
	str := strings.TrimSuffix(c.asString, "}") + ",\"replies\":["
	for idx, comment := range replies {
		str += comment.asString
		if idx != len(replies)-1 {
			str += ","
		}
	}
	str += "]}"
	return str
}

func TestThread_MarshalJSON(t *testing.T) {
	type fields struct {
		MainComment *Comment
		Replies     []*Comment
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "only main comment thread",
			fields: fields{
				MainComment: &commentHello.comment,
				Replies:     []*Comment{},
			},
			want: commentHello.AsMainCommentStr(nil),
		},
		{
			name: "one reply thread",
			fields: fields{
				MainComment: &commentHello.comment,
				Replies: []*Comment{
					&commentHello.comment,
				},
			},
			want: commentHello.AsMainCommentStr([]comment_test_Sample{commentHello}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thread := Thread{
				MainComment: tt.fields.MainComment,
				Replies:     tt.fields.Replies,
			}
			got, err := thread.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("MarshalJSON() got:\n\t`%v`\nwant:\n\t`%v`", string(got), tt.want)
			}
		})
	}
}

func TestThread_UnmarshalJSON(t *testing.T) {
	type wanted = Thread
	type args struct {
		json []byte
	}
	tests := []struct {
		name    string
		want    wanted
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "only main comment thread",
			want: wanted{
				MainComment: &commentObi.comment,
				Replies:     []*Comment{},
			},
			args: args{
				json: []byte(commentObi.AsMainCommentStr(nil)),
			},
			wantErr: false,
		},
		{
			name: "one reply thread",
			want: wanted{
				MainComment: &commentObi.comment,
				Replies:     []*Comment{&commentGrievous.comment},
			},
			args: args{
				json: []byte(commentObi.AsMainCommentStr([]comment_test_Sample{commentGrievous})),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thread := Thread{}
			if err := json.Unmarshal(tt.args.json, &thread); (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal(Thread) error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(thread, tt.want) {
				t.Errorf("missmatch on the expected value loaded.\ngot:\n\t%v\nwant:\n\t%v", thread, tt.want)
			}
		})
	}
}
