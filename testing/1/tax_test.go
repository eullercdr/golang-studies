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

// go test -fuzz=. -run=^#
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Receive %f but Expect 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Receive %f but Expect 20", result)
		}
	})
}
