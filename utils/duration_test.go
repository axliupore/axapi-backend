package utils

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		times   string
		want    time.Duration
		wantErr bool
	}{
		{
			times:   "5",
			want:    time.Hour * 5 * 24,
			wantErr: false,
		},
		{
			times:   "7",
			want:    time.Hour * 7 * 24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.times, func(t *testing.T) {
			got, err := ParseDuration(tt.times)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
