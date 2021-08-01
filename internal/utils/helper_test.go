package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloatWithFraction(t *testing.T) {
	tests := []struct {
		name     string
		in       float64
		fraction int
		want     int64
	}{
		{"", 1234.434, 2, 123443},
		{"", 1234.436, 2, 123444},
		{"", 1234.4, 2, 123440},
		{"", 1234.4, 0, 1234},
		{"", 1234.5, 0, 1235},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatWithFraction(tt.in, tt.fraction)
			require.Equal(t, tt.want, got)
		})
	}
}
