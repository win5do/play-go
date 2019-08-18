package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	foo := &Foo{
		Type:   "a",
		Object: A{A: "rua"},
	}

	bts, err := json.Marshal(foo)
	require.Nil(t, err)
	t.Logf("jsonStr = %s", bts)

	newFoo := new(Foo)
	err = json.Unmarshal(bts, newFoo)
	require.Nil(t, err)
	t.Logf("struct = %+v, object = %+v", newFoo, newFoo.Object)
}
