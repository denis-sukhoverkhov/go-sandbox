package main

import (
	"reflect"
	"testing"
)

func TestFact(t *testing.T) {
	result := Fact(5)
	if !reflect.DeepEqual(result, 120) {
		t.Errorf("Fact is not correct, got: %v, expected: %v", result, 5)
	}

	result = Fact(0)
	if !reflect.DeepEqual(result, 1) {
		t.Errorf("Fact is not correct, got: %v, expected: %v", result, 1)
	}

	result = Fact(1)
	if !reflect.DeepEqual(result, 1) {
		t.Errorf("Fact is not correct, got: %v, expected: %v", result, 1)
	}
}

func TestElapsed(t *testing.T) {
	result, err := Elapsed(30)
	if err != nil {
		t.Errorf("Expiration time")
	}
	expected := uint64(9682165104862298112)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Fact is not correct, got: %v, expected: %v", result, expected)
	}
}

func TestBusActions(t *testing.T) {

	var lst []func() error
	for i := 0; i < 100; i++ {
		lst = append(lst, WrapDummyJob)
	}
	BusActions(lst, 3, 0)
	//
	//if err != nil{
	//	t.Errorf("Expiration time",)
	//}
	//if !reflect.DeepEqual(result, 120) {
	//	t.Errorf("Fact is not correct, got: %v, expected: %v", result, 5)
	//}
}
