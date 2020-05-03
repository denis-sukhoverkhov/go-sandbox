package quizz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStruct1(t *testing.T) {
	seen := make(map[string]struct{})

	v := "qwerty"
	for _, ok := seen[v]; ok == false; {
		seen[v] = struct{}{}
	}

}

func hello(num ...int) {
	num[0] = 18
}

func TestVariadicFunction(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	assert.Equal(t, 18, i[0])

}


func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func TestVariadicFunction2(t *testing.T) {
	welcome := []string{"hello", "world"}
	assert.Equal(t, cap(welcome), 2)

	change(welcome...)
	assert.Equal(t, []string{"Go", "world"}, welcome)

}
