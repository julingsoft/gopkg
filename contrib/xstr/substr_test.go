package xstr

import "testing"

func TestSubStr(t *testing.T) {
	type args struct {
		s      string
		start  int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s:      "你好，世界！Hello, World!",
				start:  0,
				length: 5,
			},
			want: "你好，世界",
		}, {
			name: "case1",
			args: args{
				s:      "你好，世界！Hello, World!",
				start:  0,
				length: 5,
			},
			want: "你好，世界",
		}, {
			name: "case2",
			args: args{
				s:      "你好，世界！Hello, World!",
				start:  3,
				length: 5,
			},
			want: "Hello",
		}, {
			name: "case1",
			args: args{
				s:      "你好，世界！Hello, World!",
				start:  -6,
				length: 5,
			},
			want: "World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStr(tt.args.s, tt.args.start, tt.args.length); got != tt.want {
				t.Errorf("SubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
