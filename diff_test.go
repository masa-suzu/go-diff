package diff_test

import (
	"reflect"
	"testing"

	"github.com/masa-suzu/go-diff"
)

func TestDiff(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want []diff.Edit
	}{
		{
			name: "same",
			args: args{
				a: "xyz",
				b: "xyz",
			},
			want: []diff.Edit{
				{Action: 0, Value: "x"},
				{Action: 0, Value: "y"},
				{Action: 0, Value: "z"},
			},
		},
		{
			name: "added1",
			args: args{
				a: "",
				b: "x",
			},
			want: []diff.Edit{
				{Action: 1, Value: "x"},
			},
		},
		{
			name: "added2",
			args: args{
				a: "xz",
				b: "xyz",
			},
			want: []diff.Edit{
				{Action: 0, Value: "x"},
				{Action: 1, Value: "y"},
				{Action: 0, Value: "z"},
			},
		},
		{
			name: "deleted",
			args: args{
				a: "xvz",
				b: "xz",
			},
			want: []diff.Edit{
				{Action: 0, Value: "x"},
				{Action: -1, Value: "v"},
				{Action: 0, Value: "z"},
			},
		},
		{
			name: "added-and-deleted",
			args: args{
				a: "stream",
				b: "cream",
			},
			want: []diff.Edit{
				{Action: -1, Value: "s"},
				{Action: -1, Value: "t"},
				{Action: 1, Value: "c"},
				{Action: 0, Value: "r"},
				{Action: 0, Value: "e"},
				{Action: 0, Value: "a"},
				{Action: 0, Value: "m"},
			},
		},
		{
			name: "japanese",
			args: args{
				a: "hello,world!",
				b: "hello,世界!",
			},
			want: []diff.Edit{
				{Action: 0, Value: "h"},
				{Action: 0, Value: "e"},
				{Action: 0, Value: "l"},
				{Action: 0, Value: "l"},
				{Action: 0, Value: "o"},
				{Action: 0, Value: ","},
				{Action: -1, Value: "w"},
				{Action: -1, Value: "o"},
				{Action: -1, Value: "r"},
				{Action: -1, Value: "l"},
				{Action: -1, Value: "d"},
				{Action: 1, Value: "世"},
				{Action: 1, Value: "界"},
				{Action: 0, Value: "!"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // pin!
		t.Run(tt.name, func(t *testing.T) {
			got := diff.Diff(tt.args.a, tt.args.b)

			if len(tt.want) != len(got) {
				t.Errorf("len(got) = %v, len(want) %v", len(got), len(tt.want))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
