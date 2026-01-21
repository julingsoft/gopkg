package xguid

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		{
			name:    "case1",
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextID(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "case1",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				if got := NextID(); got != tt.want {
					t.Errorf("NextID() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestGetDefaultMachineID(t *testing.T) {
	tests := []struct {
		name    string
		want    uint16
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1",
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				got, err := GetDefaultMachineID()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetDefaultMachineID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("GetDefaultMachineID() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
