package main

import "testing"

// unit test for first function

func TestConverter(t *testing.T) {
	input := "Testing out our&^&## sample test for&^#& converter"
	expectedOutput := "Testing__out__our__sample__test__for__converter"

	actual, err := converter(input)
	if err != nil {
		t.Errorf("expected no error,but we got %v", err)
	}

	if actual != expectedOutput {
		t.Errorf("expected output to be %s, but we got %s", expectedOutput, actual)
	}
}

// unit test for second function

func TestCalculation(t *testing.T) {
	if calculation(1000) != 2000 {
		t.Errorf("Expected mathematical value of 1000 * 2 is 2000")
	}
}
