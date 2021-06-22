package test

import (
	"awesomeProject3/currencies"
	"testing"
)

func TestGetRate(t *testing.T) {
	FirstCurrency := "USD"
	SecondCurrency := "RUB"

	rate, err := currencies.GetRate(FirstCurrency, SecondCurrency)
	if err != nil {
		t.Errorf("function execution error")
	}

	if rate <= 0 {
		t.Errorf("error finding currency")
	}
}
