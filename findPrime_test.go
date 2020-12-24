package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	r := isPrime(15)

	if r == true {
		t.Error("isPrime(15) return true, should be false")
	}

	r = isPrime(53)

	if r == false {
		t.Error("isPrime(53) return false, should be true")
	}
}

func TestFindPrime(t *testing.T) {
	r := FindPrime(55)

	if r != 53 {
		t.Errorf("findPrime(55) return %d, should be 53", r)
	}
}
