package gekit

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
    v := "hello"
    res := ToPtr[string](v)
    assert.Equal(t, &v, res)
}
