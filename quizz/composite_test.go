package quizz

import (
	. "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) FullName() string {
	return Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) Details() string {
	return Sprintf("%s %s: bio %s", p.title, p.content, p.bio)
}

func TestComposite(t *testing.T) {
	t.Run("test1", func(t *testing.T) {

		a := author{
			firstName: "зек",
			lastName:  "зекович",
			bio:       "сиделович",
		}
		assert.Equal(t, a.FullName(), "зек зекович")

		p := post{
			title:   "title",
			content: "ctx",
			author:  a,
		}

		assert.Equal(t, "title ctx: bio сиделович", p.Details())
	})
}
