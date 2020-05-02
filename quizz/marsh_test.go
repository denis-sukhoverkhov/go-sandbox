package quizz

import (
	"fmt"
	"github.com/denis-sukhoverkhov/go-sandbox/tree/master/otus/quizz/quizzExport"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type User struct {
	Name string `key:"name" maxlength:"128"`
}

func TestMultipleTags(t *testing.T) {
	t.Run("test1", func(t *testing.T) {

		user := User{Name: "Сибур"}
		st := reflect.TypeOf(user)
		field := st.Field(0)
		assert.Equal(t, "name", field.Tag.Get("key"))
		assert.Equal(t, "128", field.Tag.Get("maxlength"))
	})
}

func TestDefer(t *testing.T) {
	t.Run("defer", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			defer fmt.Println(i)
		}
	})
}

func TestMapOrder(t *testing.T) {
	t.Run("map order", func(t *testing.T) {
		tagViews := map[string]int{
			"unix":   0,
			"python": 1,
			"go":     2,
			"golang": 3,
			"devops": 4,
			"gc":     5,
		}
		for key, val := range tagViews {
			fmt.Println("There are", val, "views for", key)
		}
	})
}

func TestExport(t *testing.T) {
	assert.Equal(t, "", quizzExport.BigBro)
	assert.Equal(t, "", quizzExport.Долли)
}
