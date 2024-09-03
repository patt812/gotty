package utils_test

import (
	"gotty/test/utils"
	"testing"
)

func TestRemoveANSISequences(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "No ANSI sequences",
			input: "This is a normal string",
			want:  "This is a normal string",
		},
		{
			name:  "String with Carriage Return not removed",
			input: "This is a string with \rCarriage Return",
			want:  "This is a string with \rCarriage Return",
		},
		{
			name:  "String with ANSI sequences",
			input: "\x1b[31mThis is red text\x1b[0m",
			want:  "This is red text",
		},
		{
			name:  "String with multiple ANSI sequences",
			input: "\x1b[31mRed\x1b[0m and \x1b[32mGreen\x1b[0m",
			want:  "Red and Green",
		},
		{
			name:  "String with ANSI and control characters",
			input: "Start\x1b[31mRed\x1b[0m\rEnd",
			want:  "StartRed\rEnd",
		},
		{
			name:  "String with multiple cursor movements",
			input: "\x1b[1;0H\r\x1b[Kこんにちは\x1b[2;0H\r\x1b[Kこんにちは\x1b[5;0H\r\x1b[KAccuracy: --- WPM: 0.00 (Current: 0.00)",
			want:  "\rこんにちは\rこんにちは\rAccuracy: --- WPM: 0.00 (Current: 0.00)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := utils.RemoveANSISequences(tt.input)
			if got != tt.want {
				t.Errorf("RemoveANSISequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
