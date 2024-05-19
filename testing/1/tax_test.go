package tax

import "testing"

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
