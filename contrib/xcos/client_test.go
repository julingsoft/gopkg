package xcos_test

import (
	"testing"
)

func TestGetCosKeyFromURLSafe(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rawURL  string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "valid URL",
			rawURL:  "https://bucket.cos.ap-shanghai.myqcloud.com/path/to/file?a=asdg&b=asdfg",
			want:    "path/to/file",
			wantErr: false,
		},
		{
			name:    "invalid URL",
			rawURL:  "https://bucket.cos.ap-shanghai.myqcloud.com/path/to/file",
			want:    "path/to/file",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCosKeyFromURLSafe(tt.rawURL)
			if got != tt.want {
				t.Errorf("GetCosKeyFromURLSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
