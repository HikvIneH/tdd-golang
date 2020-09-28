package tdd_golang_test

import (
	"fmt"
	"testing"

	tdd_golang "github.com/supriyadi687/tdd-golang"
)

func TestShouldReturn0GivenEmptyString(t *testing.T) {
	res, _ := tdd_golang.Add("")
	if res != 0 {
		t.Errorf("expected to be 0 got %d", res)
	}
}

func TestShouldReturnOneGivenOneIntString(t *testing.T) {
	res, _ := tdd_golang.Add("1")
	if res != 1 {
		t.Errorf("expected to be 1 got %d", res)
	}
}

func TestShouldReturnSumGivenSumIntString(t *testing.T) {
	res, _ := tdd_golang.Add("1,2")
	if res != 3 {
		t.Errorf("expected to be 3 got %d", res)
	}
}

func TestShouldReturnSumGivenSumIntStringWithNewLine(t *testing.T) {
	res, _ := tdd_golang.Add("1,2\n3")
	if res != 6 {
		t.Errorf("expected to be 6 got %d", res)
	}
}

func TestShouldReturnSumGivenDifferentDelimiters(t *testing.T) {
	res, _ := tdd_golang.Add("//;\n1;3")
	fmt.Print(res)
	if res != 4 {
		t.Errorf("expected to be 4 got %d", res)
	}
}

func TestShouldHandleNegativeNumber(t *testing.T) {
	res, _ := tdd_golang.Add("1,4,-1")
	fmt.Print(res)
	t.Errorf("negatives not allowed: %d", res)
}
