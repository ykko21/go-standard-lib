package main

import (
	"testing"
)

//go test -v

var tests = []struct {
	name          string
	dividend      float32
	divisor       float32
	expected      float32
	isErrExpected bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		v, err := divide(tt.dividend, tt.divisor)
		if tt.isErrExpected {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one")
			}
		}

		if v != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, v)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
