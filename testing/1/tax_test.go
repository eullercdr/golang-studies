package tax

import "testing"

// go test -bench=.
// go test -html=coverage.out
func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 25.0
	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{500.0, 25.0},
		{1000.0, 5.0},
		{1500.0, 5.0},
	}
	for _, test := range tests {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}

// go test -bench=.
// go test -bench=. -run=^#
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// go test -bench=. -run=^# -count=10 -benchtime=3s
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
