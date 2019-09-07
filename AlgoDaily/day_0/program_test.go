package main

import "testing"

func TestReverse(t *testing.T) {
	expected := "ThisIsTestWith@#SpecailCharahters"

	result := reverse("sretharahCliacepS#@htiWtseTsIsihT")

	if result != expected {
		t.Errorf("reverse() failed expected %v, got %v",expected, result)
	}
}

func TestOptimizedReverse(t *testing.T) {
	expected := "ThisIsTestWith@#SpecailCharahters"

	result := optimizedReverse("sretharahCliacepS#@htiWtseTsIsihT")

	if result != expected {
		t.Errorf("optimizedReverse() failed expected %v, got %v",expected, result)
	}
}