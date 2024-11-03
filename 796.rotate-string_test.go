package main

import (
	"testing"
)

func TestRotateString(t *testing.T) {
	tests := []struct {
		s    string
		goal string
		want bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
		{"a", "a", true},
		{"ab", "ba", true},
		{"ab", "ab", true},
	}

	for _, tt := range tests {
		t.Run(tt.s+"_"+tt.goal, func(t *testing.T) {
			if got := rotateString(tt.s, tt.goal); got != tt.want {
				t.Errorf("rotateString(%v, %v) = %v; want %v", tt.s, tt.goal, got, tt.want)
			}
		})
	}
}

func TestRotateString2(t *testing.T) {
	tests := []struct {
		s    string
		goal string
		want bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
		{"a", "a", true},
		{"ab", "ba", true},
		{"ab", "ab", true},
	}

	for _, tt := range tests {
		t.Run(tt.s+"_"+tt.goal, func(t *testing.T) {
			if got := rotateString2(tt.s, tt.goal); got != tt.want {
				t.Errorf("rotateString2(%v, %v) = %v; want %v", tt.s, tt.goal, got, tt.want)
			}
		})
	}
}

func BenchmarkRotateString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateString("abcde", "cdeab")
	}
}

func BenchmarkRotateString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateString2("abcde", "cdeab")
	}
}
