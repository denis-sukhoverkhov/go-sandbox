package quizz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"unsafe"
)

const Pi = math.Pi

const (
	a = iota + 1
	b
	c
)

const (
	d, e, f = iota, iota, iota
)

func test1() (int, int, int, int, int, int) {
	return a, b, c, d, e, f
}

var r, t, y = 4, "sfasdas", true

func TestIota(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		a, b, c, d, e, f := test1()

		assert.Equal(t, []int{a, b, c, d, e, f}, []int{1, 2, 3, 0, 0, 0})

	})
	fmt.Println(r, t, y)
}

func TestSizeOf(t *testing.T) {

	s := struct {
		A float32
		B string
	}{0, "go"}

	ss := struct{}{}

	fmt.Printf("%T, %d\n", s, unsafe.Sizeof(ss))
}

func testArr(arr [2]int) {
	arr[0] = 5
}

func testArrP(arr *[2]int) {
	arr[0] = 5
}

func testArrSl(arr []int) {
	arr[0] = 7
}

func TestArray(t *testing.T) {

	arr := [2]int{3, 4}
	testArr(arr)
	assert.Equal(t, arr, [2]int{3, 4})

	testArrP(&arr)
	assert.Equal(t, arr, [2]int{5, 4})

	sl := []int{9, 9}
	testArrSl(sl)
	assert.Equal(t, sl, []int{7, 9})
}
