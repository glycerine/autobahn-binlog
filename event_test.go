package binlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	someEventHeader = []byte{0x52, 0x52, 0xe9, 0x53, 0xf, 0x88, 0xf3, 0x0, 0x0, 0x66, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x35, 0x2e, 0x31, 0x2e, 0x36, 0x33, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2d, 0x6c, 0x6f, 0x67, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1b, 0x38, 0xd, 0x0, 0x8, 0x0, 0x12, 0x0, 0x4, 0x4, 0x4, 0x4, 0x12, 0x0, 0x0, 0x53, 0x0, 0x4, 0x1a, 0x8, 0x0, 0x0, 0x0, 0x8, 0x8, 0x8, 0x2}
)

func TestEventErrorContainsError(t *testing.T) {
	h, _ := NewEventHeader(someEventHeader)
	s := "test error"
	e := &EventError{h, s, []byte{}}

	assert.Equal(t, e.Error(), s)
}