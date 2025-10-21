package xslice

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				slice: []string{"a", "b", "c", "a", "b", "c"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
