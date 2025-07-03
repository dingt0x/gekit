package slice

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestTransform(t *testing.T) {
    a := "hello"
    b := "world"
    c := "!"

    v := []string{a, b, c}
    res := Transform(v, func(i int, s string) *string {
        return &s
    })

    assert.Equal(t, []*string{&a, &b, &c}, res)

}
