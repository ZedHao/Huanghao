package design_pattern_test

import (
    "testing"


    "github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
    assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
                b.Errorf("test fail")
            }
        }
    })
}