package chapter_6

import "testing"

func TestThreeValues(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
		want2 float32
	}{
		{"0", 5, 6, 7.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ThreeValues()
			if got != tt.want {
				t.Errorf("ThreeValues() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ThreeValues() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ThreeValues() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
