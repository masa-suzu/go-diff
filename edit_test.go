package diff_test

import (
	"testing"

	"github.com/masa-suzu/go-diff"
)

func TestEditString(t *testing.T) {
	tests := []struct {
		name string
		arg  diff.Edit
		want string
	}{
		{
			name: "add",
			want: "+xx",
			arg:  diff.Edit{Action: 1, Value: "xx"},
		},
		{
			name: "delete",
			want: "-yy",
			arg:  diff.Edit{Action: -1, Value: "yy"},
		},
		{
			name: "unChange",
			want: "x",
			arg:  diff.Edit{Action: 0, Value: "x"},
		},
	}
	for _, tt := range tests {
		tt := tt // pin!
		t.Run(tt.name, func(t *testing.T) {
			e := tt.arg
			if got := e.String(); got != tt.want {
				t.Errorf("Edit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEditScript_String(t *testing.T) {
	tests := []struct {
		name string
		es   diff.EditScript
		want string
	}{
		{
			name: "unChange-add-delete",
			es: diff.EditScript{
				diff.Edit{Action: 0, Value: "qaz"},
				diff.Edit{Action: 1, Value: "wsx"},
				diff.Edit{Action: -1, Value: "edc"},
			},
			want: "qaz\n+wsx\n-edc",
		},
	}
	for _, tt := range tests {
		tt := tt // pin!
		t.Run(tt.name, func(t *testing.T) {
			if got := diff.String(tt.es); got != tt.want {
				d := diff.Diff(tt.want, got)
				t.Errorf("\nToString(tt.es):\n\t%#v\nwant:\n\t%#v\nDiff:\n%v", got, tt.want, diff.String(d))
			}
		})
	}
}
