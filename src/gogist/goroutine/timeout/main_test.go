package timeout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	var r int
	f := func() {
		r = add(1, 2)
	}
	err := timeoutWrap(f)
	assert.Error(t, err)
	assert.Zero(t, r)
}
