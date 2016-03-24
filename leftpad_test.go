package leftpad

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftPad(t *testing.T) {
	assert.Equal(t, "  foo ", LeftPad("foo ", 6, " "))
	assert.Equal(t, "foo", LeftPad("foo", 3, " "))
}
