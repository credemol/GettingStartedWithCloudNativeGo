package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Cloud Native Go","Author":"M.-L. Reimer","ISBN":"0123456789"}`,
		string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
}
