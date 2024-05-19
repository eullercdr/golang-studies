package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxAndSaveShouldNotReturnAError(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	err := CalculateTaxAndSAve(1000.0, repository)
	assert.Nil(t, err)
	repository.AssertExpectations(t)
}
func TestCalculateTaxAndSaveShouldReturnAError(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 0.0).Return(errors.New("amount should not 0"))
	err := CalculateTaxAndSAve(0, repository)
	assert.NotNil(t, err, "amount should not 0")
	repository.AssertExpectations(t)
}
func TestCalculateTaxAndSaveShouldReturnAErrorShouldCallOneTime(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 0.0).Return(errors.New("amount should not 0")).Once()
	err := CalculateTaxAndSAve(0, repository)
	assert.NotNil(t, err, "amount should not 0")
	repository.AssertExpectations(t)
	repository.AssertCalled(t, "SaveTax", 0.0)
}
