package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := [][]string{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for _, v := range tests {
		result := unpack(v[0])

		shouldBe := v[1]
		if result != shouldBe {
			t.Errorf("Result is not correct, got: %s, want: %s", result, shouldBe)
		}
	}

}
