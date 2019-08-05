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
				{Action: 0, Value: "xyz"},
			},
		},
		{
			name: "added2",
			args: args{
				a: "",
				b: "y\nz",
			},
			want: []diff.Edit{
				{Action: -1, Value: ""},
				{Action: 1, Value: "y"},
				{Action: 1, Value: "z"},
			},
		},
		{
			name: "deleted",
			args: args{
				a: "x\ny",
				b: "",
			},
			want: []diff.Edit{
				{Action: -1, Value: "x"},
				{Action: -1, Value: "y"},
				{Action: 1, Value: ""},
			},
		},
		{
			name: "add-delete",
			args: args{
				a: `1
2
4
5
`,
				b: `1
2
3
4
`,
			},
			want: []diff.Edit{
				{Action: 0, Value: "1"},
				{Action: 0, Value: "2"},
				{Action: 1, Value: "3"},
				{Action: 0, Value: "4"},
				{Action: -1, Value: "5"},
				{Action: 0, Value: ""},
			},
		},

		{
			name: "add-deleted-replace",
			args: args{
				a: `x := "hello"
w := z`,
				b: `x := "hello"
y := "world"
w := x`,
			},
			want: []diff.Edit{
				{Action: 0, Value: `x := "hello"`},
				{Action: -1, Value: "w := z"},
				{Action: 1, Value: `y := "world"`},
				{Action: 1, Value: "w := x"},
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
				t.Errorf("\nDiff()\n\t%v\nwant\n\t%v", got, tt.want)
			}
		})
	}
}
