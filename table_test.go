package ztable

import (
	"fmt"
	"testing"
)

func TestNewTable(t *testing.T) {
	z := NewTable()

	if z == nil {
		t.Errorf("ZTable is nil")
	}
}

func TestTableSize(t *testing.T) {
	z := NewTable()

	size := z.Len()
	want := 410
	if size != want {
		t.Errorf("Table Size = %d; wanted %d", size, want)
	}
}

func TestGetFuzzy(t *testing.T) {
	z := NewTable()

	given := 0.975
	scores := z.GetFuzzy(given)
	size := len(scores)
	if len(scores) != 3 {
		t.Errorf("GetFuzzy(%.02f); wrong size (%d), wanted %d", given, size, 3)
		return
	}

	if scores[0].Score != 1.95 {
		t.Errorf("GetFuzzy(%.02f) = %.2f, wanted %.2f", given, scores[0].Score, 1.95)
	}

	if scores[1].Score != 1.96 {
		t.Errorf("GetFuzzy(%.02f) = %.2f, wanted %.2f", given, scores[1].Score, 1.96)
	}

	if scores[2].Score != 1.97 {
		t.Errorf("GetFuzzy(%.02f) = %.2f, wanted %.2f", given, scores[2].Score, 1.97)
	}
}

func TestGetScore(t *testing.T) {
	z := NewTable()

	tests := []struct {
		given, want float64
	}{
		{0.50, 0.00},
		{0.975, 1.96},
		{1, 4.09},
		{90.5, 4.09},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.02f,%.02f", tt.given, tt.want)
		t.Run(testname, func(t *testing.T) {
			s := z.GetScore(tt.given)
			if tt.want != s {
				t.Errorf("GetScore(%.02f) = %.02f; wanted %.02f", tt.given, s, tt.want)
			}
		})
	}
}

func TestGetProbability(t *testing.T) {
	z := NewTable()

	tests := []struct {
		given, want float64
	}{
		{1.23, 0.8906514476},
		{1.63, 0.9484492515},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.02f,%.02f", tt.given, tt.want)
		t.Run(testname, func(t *testing.T) {
			s := z.GetProbability(tt.given)
			if tt.want != s {
				fmt.Println(tt.want)
				fmt.Println(s)
				t.Errorf("GetProbability(%.02f) = %.02f; wanted %.02f", tt.given, s, tt.want)
			}
		})
	}
}
