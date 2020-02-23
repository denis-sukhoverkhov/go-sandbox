package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefCounter(t *testing.T) {
	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the 
		industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and 
		scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into 
		electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of 
		Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus 
		PageMaker including versions of Lorem Ipsum.`

	result := DefCounter(text)
	fmt.Println(result)

	shouldBe := PairList{
		Pair{"the", 6},
		Pair{"of", 4},
		Pair{"Lorem", 4},
		Pair{"Ipsum", 3},
		Pair{"and", 3},
		Pair{"text", 2},
		Pair{"It", 2},
		Pair{"type", 2},
		Pair{"with", 2},
		Pair{"a", 2},
	}

	if !reflect.DeepEqual(result, shouldBe) {
		t.Errorf("Result is not correct")
	}
}
