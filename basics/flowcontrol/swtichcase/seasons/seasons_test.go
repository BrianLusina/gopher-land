package seasons

import (
	"testing"
)

func TestSeasons(t *testing.T) {
	tests := []struct {
		name  string
		month int
		want  string
	}{
		{
			name:  "January",
			month: 1,
			want:  "Winter",
		},
		{
			name:  "February",
			month: 2,
			want:  "Winter",
		},
		{
			name:  "March",
			month: 3,
			want:  "Spring",
		},
		{
			name:  "April",
			month: 4,
			want:  "Spring",
		},
		{
			name:  "May",
			month: 5,
			want:  "Spring",
		},
		{
			name:  "June",
			month: 6,
			want:  "Summer",
		},
		{
			name:  "July",
			month: 7,
			want:  "Summer",
		},
		{
			name:  "August",
			month: 8,
			want:  "Summer",
		},
		{
			name:  "September",
			month: 9,
			want:  "Autumn",
		},
		{
			name:  "October",
			month: 10,
			want:  "Autumn",
		},
		{
			name:  "November",
			month: 11,
			want:  "Autumn",
		},
		{
			name:  "December",
			month: 12,
			want:  "Winter",
		},
		{
			name:  "Unknown",
			month: 13,
			want:  "Season unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Seasons(tt.month); got != tt.want {
				t.Errorf("Seasons() = %v, want %v", got, tt.want)
			}
		})
	}
}
